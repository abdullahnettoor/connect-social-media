package server

import (
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/server/handlers"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/server/routes"
	"github.com/gin-gonic/gin"
)

type ServeHttp struct {
	server *gin.Engine
}

func NewServeHttp(
	userHandler *handlers.UserHandler,
	adminHandler *handlers.AdminHandler,
	postHandler *handlers.PostHandler,
	commentHandler *handlers.CommentHandler,
	chatHandler *handlers.ChatHandler,
	wsHandler *handlers.WebSocketConnection,
) *ServeHttp {
	server := gin.New()
	server.Use(gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/browser"}}))

	landingPage := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>ConnectR - Social Media Backend</title>
    <style>
        :root {
            --primary-color: #3ECF8E;
            --background-color: #1A1A1A;
            --text-color: #E0E0E0;
            --card-background: #2A2A2A;
            --muted-color: #A0A0A0;
            --border-color: #3A3A3A;
        }

        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
            line-height: 1.6;
            color: var(--text-color);
            background-color: var(--background-color);
        }

        .container {
            max-width: 1200px;
            margin: 0 auto;
            padding: 0 20px;
        }

        header {
            padding: 20px 0;
        }

        .logo {
            font-size: 24px;
            font-weight: bold;
            color: var(--primary-color);
        }

        .hero {
            text-align: center;
            padding: 80px 0;
        }

        h1 {
            font-size: 48px;
            margin-bottom: 20px;
        }

        p {
            font-size: 18px;
            color: var(--muted-color);
            margin-bottom: 40px;
        }

        .btn {
            display: inline-block;
            padding: 12px 24px;
            margin: 0 10px 10px;
            text-decoration: none;
            font-weight: 600;
            border-radius: 4px;
            transition: all 0.3s ease;
        }

        .btn-primary {
            background-color: var(--primary-color);
            color: var(--background-color);
        }

        .btn-secondary {
            background-color: transparent;
            color: var(--primary-color);
            border: 2px solid var(--primary-color);
        }

        .btn:hover {
            opacity: 0.9;
            transform: translateY(-2px);
        }

        .features {
            display: flex;
            flex-wrap: wrap;
            justify-content: space-between;
            margin-top: 60px;
        }

        .feature {
            flex-basis: calc(33.333% - 20px);
            margin-bottom: 40px;
            background-color: var(--card-background);
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        }

        .feature h3 {
            margin-bottom: 10px;
            color: var(--primary-color);
        }

        .feature p {
            font-size: 16px;
        }

        footer {
            text-align: center;
            padding: 20px 0;
            margin-top: 60px;
            border-top: 1px solid var(--border-color);
            color: var(--muted-color);
        }

        @media (max-width: 768px) {
            .feature {
                flex-basis: 100%;
            }
        }
    </style>
</head>
<body>
    <header>
        <div class="container">
            <div class="logo">ConnectR</div>
        </div>
    </header>

    <main>
        <section class="hero">
            <div class="container">
                <h1>ConnectR - Social Media Backend</h1>
                <p>A robust social media backend application built with Go, featuring Neo4j graph database for social connections and Kafka for message queuing.</p>
                <a href="https://github.com/abdullahnettoor/connect-social-media" class="btn btn-primary">View on GitHub</a>
                <a href="https://www.postman.com/universal-shadow-841176/workspace/connectr-socialmedia" class="btn btn-secondary">API Docs in Postman</a>
            </div>
        </section>

        <section class="features">
            <div class="container">
                <div class="feature">
                    <h3>User Authentication</h3>
                    <p>Secure user authentication and authorization system.</p>
                </div>
                <div class="feature">
                    <h3>Real-time Messaging</h3>
                    <p>WebSocket-based real-time communication.</p>
                </div>
                <div class="feature">
                    <h3>Graph Database</h3>
                    <p>Neo4j for efficient social connections management.</p>
                </div>
                <div class="feature">
                    <h3>Message Queuing</h3>
                    <p>Kafka for reliable message queuing and processing.</p>
                </div>
                <div class="feature">
                    <h3>Cloud Storage</h3>
                    <p>Cloudinary integration for media storage and management.</p>
                </div>
                <div class="feature">
                    <h3>Containerization</h3>
                    <p>Docker and Docker Compose for easy deployment and scaling.</p>
                </div>
            </div>
        </section>
    </main>

    <footer>
        <div class="container">
            <p>&copy; 2024 ConnectR. All rights reserved.</p>
        </div>
    </footer>
</body>
</html>
	`

	server.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, landingPage, nil)
	})

	routes.SetupUserRoutes(server, userHandler, postHandler, commentHandler, chatHandler, wsHandler)
	routes.SetupAdminRoutes(server, adminHandler)

	return &ServeHttp{server}
}

func (s *ServeHttp) Start() {
	s.server.Run(":9000")
}
