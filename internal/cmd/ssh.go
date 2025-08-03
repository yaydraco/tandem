package cmd

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/keygen"
	charmLog "github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/spf13/cobra"
	"github.com/yaydraco/tandem/internal/app"
	"github.com/yaydraco/tandem/internal/config"
	"github.com/yaydraco/tandem/internal/db"
	internalLogging "github.com/yaydraco/tandem/internal/logging"
	"github.com/yaydraco/tandem/internal/tui"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Start tandem SSH server for remote access",
	Long: `Start an SSH server that allows remote users to connect to tandem's TUI interface.
This enables multi-user access to tandem's penetration testing capabilities over SSH.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetInt("port")
		keyPath, _ := cmd.Flags().GetString("key-path")
		debug, _ := cmd.Flags().GetBool("debug")
		cwd, _ := cmd.Flags().GetString("cwd")

		if cwd != "" {
			err := os.Chdir(cwd)
			if err != nil {
				return fmt.Errorf("failed to change directory: %v", err)
			}
		}
		if cwd == "" {
			c, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("failed to get current working directory: %v", err)
			}
			cwd = c
		}

		// Load the config
		_, err := config.Load(cwd, debug)
		if err != nil {
			return err
		}

		// Connect DB, this will also run migrations
		conn, err := db.Connect()
		if err != nil {
			return err
		}

		// Create main context for the application
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		app, err := app.New(ctx, conn)
		if err != nil {
			internalLogging.Error("Failed to create app: %v", err)
			return err
		}

		// Generate or load SSH host key
		k, err := keygen.New(keyPath, keygen.WithKeyType(keygen.Ed25519))
		if err != nil {
			return fmt.Errorf("failed to create or load SSH key: %v", err)
		}

		// Configure SSH server
		s, err := wish.NewServer(
			wish.WithAddress(net.JoinHostPort(host, fmt.Sprintf("%d", port))),
			wish.WithHostKeyPEM(k.RawPrivateKey()),
			wish.WithMiddleware(
				bubbletea.Middleware(func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
					// Create a new TUI instance for each SSH session
					return tui.New(app), []tea.ProgramOption{tea.WithAltScreen()}
				}),
				logging.Middleware(),
			),
		)
		if err != nil {
			return fmt.Errorf("failed to create SSH server: %v", err)
		}

		// Setup graceful shutdown
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		// Start server in a goroutine
		go func() {
			internalLogging.Info("Starting SSH server on %s:%d", host, port)
			if err := s.ListenAndServe(); err != nil && err != ssh.ErrServerClosed {
				charmLog.Error("SSH server failed", "error", err)
			}
		}()

		// Wait for shutdown signal
		<-done
		internalLogging.Info("Shutting down SSH server...")

		// Create shutdown context with timeout
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer shutdownCancel()

		// Shutdown the server
		if err := s.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("SSH server forced to shutdown: %v", err)
		}

		internalLogging.Info("SSH server stopped")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)

	// SSH server configuration flags
	sshCmd.Flags().String("host", "localhost", "Host to bind the SSH server to")
	sshCmd.Flags().Int("port", 2222, "Port to bind the SSH server to")
	sshCmd.Flags().String("key-path", ".ssh/tandem_host_key", "Path to SSH host key (will be generated if doesn't exist)")
	sshCmd.Flags().BoolP("debug", "d", false, "Enable debug logging")
	sshCmd.Flags().StringP("cwd", "c", "", "Current working directory")
}