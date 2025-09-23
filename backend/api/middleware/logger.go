package middlewares

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/labstack/echo/v4"
)

func styleStatus(status int) string {
	var style lipgloss.Style
	switch {
	case status >= 200 && status < 300:
		style = lipgloss.NewStyle().Background(lipgloss.Color("#33c100ff")).Bold(true)
	case status >= 400 && status < 500:
		style = lipgloss.NewStyle().Background(lipgloss.Color("#ffae00ff")).Bold(true)
	case status >= 500:
		style = lipgloss.NewStyle().Background(lipgloss.Color("#eb0000ff")).Bold(true)
	default:
		style = lipgloss.NewStyle().Bold(true)
	}

	return style.Render(fmt.Sprintf(" %d ", status))
}

func styleMethod(method string) string {
	var style lipgloss.Style
	switch method {
	case "GET":
		style = lipgloss.NewStyle().Background(lipgloss.Color("#0084ffff")).Bold(true)
	case "POST":
		style = lipgloss.NewStyle().Background(lipgloss.Color("#00b828ff")).Bold(true)
	case "PUT":
		style = lipgloss.NewStyle().Background(lipgloss.Color("#ffa200ff")).Bold(true)
	case "DELETE":
		style = lipgloss.NewStyle().Background(lipgloss.Color("#ff00d4")).Bold(true)
	case "PATCH":
		style = lipgloss.NewStyle().Background(lipgloss.Color("#ff6a00")).Bold(true)
	case "HEAD":
		style = lipgloss.NewStyle().Background(lipgloss.Color("#d400ff")).Bold(true)
	case "OPTIONS":
		style = lipgloss.NewStyle().Background(lipgloss.Color("#ff006a")).Bold(true)
	default:
		style = lipgloss.NewStyle().Bold(true)
	}

	return style.Render(fmt.Sprintf(" %-7s ", method))
}

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		res := c.Response()

		start := time.Now()
		err := next(c)
		if err != nil {
			c.Error(err)
		}

		dateTime := time.Now().Format("02/01/2006 - 15:04:05")
		status := styleStatus(res.Status)
		latency := time.Since(start)
		ip := c.RealIP()
		method := styleMethod(req.Method)
		path := req.URL.Path

		fmt.Printf("%s |%s| %13v | %15s |%s  %s\n", dateTime, status, latency, ip, method, path)
		return nil
	}
}
