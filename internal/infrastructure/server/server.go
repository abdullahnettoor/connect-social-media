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
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>ConnectR - Social Media Backend</title>
    <style>
        :root {
            --primary-color: #0070f3;
            --background-color: #000;
            --text-color: #fff;
            --card-background: #111;
            --muted-color: #888;
            --border-color: #333;
            --gradient-start: #7928ca;
            --gradient-end: #ff0080;
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
            color: var(--text-color);
        }

        main {
            position: relative;
            z-index: 1;
        }

        .hero {
            text-align: center;
            padding: 120px 0 180px;
            background: linear-gradient(135deg, var(--gradient-start), var(--gradient-end));
            position: relative;
            overflow: hidden;
        }

        .hero::after {
            content: '';
            position: absolute;
            bottom: 0;
            left: 0;
            right: 0;
            height: 100px;
            background: var(--background-color);
            transform: skewY(-3deg);
            transform-origin: 100%;
        }

        h1 {
            font-size: 64px;
            margin-bottom: 20px;
            line-height: 1.2;
            background: linear-gradient(to right, #fff, #f0f0f0);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }

        p {
            font-size: 20px;
            color: var(--text-color);
            margin-bottom: 40px;
            max-width: 600px;
            margin-left: auto;
            margin-right: auto;
        }

        .btn {
            display: inline-block;
            padding: 12px 24px;
            margin: 0 10px 10px;
            text-decoration: none;
            font-weight: 600;
            border-radius: 5px;
            transition: all 0.3s ease;
            font-size: 16px;
        }

        .btn-primary {
            background-color: var(--primary-color);
            color: var(--text-color);
        }

        .btn-secondary {
            background-color: transparent;
            color: var(--text-color);
            border: 2px solid var(--text-color);
        }

        .btn:hover {
            transform: translateY(-2px);
            box-shadow: 0 4px 6px rgba(50, 50, 93, 0.11), 0 1px 3px rgba(0, 0, 0, 0.08);
        }

        .features {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
            gap: 30px;
            margin-top: 60px;
            padding: 0 20px;
            position: relative;
            z-index: 2;
        }

        .feature {
            background-color: var(--card-background);
            padding: 30px;
            border-radius: 8px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
            transition: all 0.3s ease;
        }

        .feature:hover {
            transform: translateY(-5px);
            box-shadow: 0 8px 12px rgba(0, 0, 0, 0.2);
        }

        .feature h3 {
            margin-bottom: 15px;
            color: var(--primary-color);
            font-size: 24px;
        }

        .feature p {
            font-size: 16px;
            color: var(--muted-color);
        }

        footer {
            text-align: center;
            padding: 40px 0;
            margin-top: 80px;
            border-top: 1px solid var(--border-color);
            color: var(--muted-color);
        }

        .maker {
            text-align: center;
            padding: 60px 0;
            background-color: var(--card-background);
            margin-top: 80px;
        }

        .maker-content {
            max-width: 600px;
            margin: 0 auto;
        }

        .maker h2 {
            font-size: 32px;
            margin-bottom: 20px;
            background: linear-gradient(to right, var(--primary-color), var(--gradient-end));
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }

        .maker-link {
            color: var(--primary-color);
            text-decoration: none;
            font-weight: 500;
            transition: color 0.3s ease;
        }

        .maker-link:hover {
            color: var(--gradient-end);
        }

        @media (max-width: 768px) {
            h1 {
                font-size: 48px;
            }

            .hero {
                padding: 80px 0;
            }

            .features {
                margin-top: -40px;
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
                <a href="https://www.postman.com/abdullahnettoor/workspace/connectr-socialmedia" class="btn btn-secondary">API Docs in Postman</a>
            </div>
        </section>

        <section class="features">
            <div class="feature">
                <h3>User Authentication</h3>
                <p>Secure user authentication and authorization system to protect user data and ensure privacy.</p>
            </div>
            <div class="feature">
                <h3>Real-time Messaging</h3>
                <p>WebSocket-based real-time communication for instant messaging and live updates.</p>
            </div>
            <div class="feature">
                <h3>Graph Database</h3>
                <p>Neo4j for efficient social connections management, enabling complex relationship queries.</p>
            </div>
            <div class="feature">
                <h3>Message Queuing</h3>
                <p>Kafka for reliable message queuing and processing, ensuring scalability and fault tolerance.</p>
            </div>
            <div class="feature">
                <h3>Cloud Storage</h3>
                <p>Cloudinary integration for seamless media storage and management in the cloud.</p>
            </div>
            <div class="feature">
                <h3>Containerization</h3>
                <p>Docker and Docker Compose for easy deployment, scaling, and management of the application.</p>
            </div>
        </section>
        <section class="maker">
            <div class="container">
                <div class="maker-content">
                    <h2>Meet the Maker</h2>
                    <p>Built with ❤️ by <a href="https://github.com/abdullahnettoor" class="maker-link">Abdullah Nettoor</a></p>
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
	`

	server.GET("/", func(ctx *gin.Context) {
        ctx.Header("Content-Type", "text/html")
		ctx.String(200, "<!DOCTYPE html>\n<html>%s</html>", landingPage)
	})

	routes.SetupUserRoutes(server, userHandler, postHandler, commentHandler, chatHandler, wsHandler)
	routes.SetupAdminRoutes(server, adminHandler)

	return &ServeHttp{server}
}

func (s *ServeHttp) Start() {
	s.server.Run(":9000")
}
