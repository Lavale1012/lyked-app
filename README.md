# 💡 Lyked

> **Your Personal Media Vault** - Save, organize, and rediscover your favorite content from across the web.

[![React Native](https://img.shields.io/badge/React%20Native-61DAFB?style=for-the-badge&logo=react&logoColor=black)](https://reactnative.dev/)
[![Expo](https://img.shields.io/badge/Expo-000020?style=for-the-badge&logo=expo&logoColor=white)](https://expo.dev/)
[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![MongoDB](https://img.shields.io/badge/MongoDB-47A248?style=for-the-badge&logo=mongodb&logoColor=white)](https://mongodb.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=for-the-badge&logo=postgresql&logoColor=white)](https://postgresql.org/)

**Lyked** is a mobile application that empowers users to save and organize their favorite videos and images from various social media platforms — such as TikTok, Instagram, YouTube, and Pinterest — in a private, personalized media vault.

---

## 🧩 The Problem

In the age of infinite scrolling, users often come across inspiring, funny, emotional, or educational content — only to lose it moments later. Social media apps aren't designed for long-term media curation. Users resort to scattered bookmarks, screenshots, or saved folders that are platform-locked and difficult to manage.

## ✅ The Solution

**Lyked** gives users one unified space to:

- 📎 **Collect links** to favorite images and videos from across the internet
- 🏷️ **Categorize and tag** media for easy filtering and discovery
- 📱 **View all saved content** in a beautiful, scrollable visual library
- 🗂️ **Organize content** without relying on platform-specific saved features
- 🔍 **Search and filter** through your curated collection effortlessly

> *Future versions will include onboard AI that learns from stored media and recommends similar videos and images from the internet — tailoring discovery to your true tastes.*

---

## 📱 Screenshots

*Coming Soon - Screenshots will be added as the UI is finalized*

---

## 🚀 Features

### ✅ Current (MVP)
- **Cross-Platform Mobile App** built with React Native + Expo
- **Media Link Storage** with automatic thumbnail generation
- **Folder Organization** for categorizing saved content
- **Tag Management** for flexible content organization
- **Visual Grid Layout** for browsing saved media
- **REST API Backend** with secure data storage

### 🔄 In Progress
- **User Authentication** via Clerk (email & social login)
- **Enhanced Create Form** with URL validation and preview
- **Advanced Search & Filtering** by tags, folders, and content type
- **Error Boundaries** and offline capability

### ⏳ Planned Features
- **AI-Powered Recommendations** based on saved content
- **Full-Text Search** across titles, descriptions, and tags
- **Web Browser Extension** for one-click saving
- **Social Features** - Share collections with friends
- **Advanced Analytics** - Content insights and trends

---

## 🏗️ Architecture

### 📱 Frontend (`/mobile`)
- **React Native + Expo** - Cross-platform mobile development
- **NativeWind** - Tailwind CSS styling for React Native
- **React Navigation** - Tab and stack navigation
- **Context API** - User management and state
- **TypeScript** - Type-safe development

### 🚀 Backend (`/backend`)
- **Go + Gin** - High-performance REST API server
- **Dual Database Architecture**:
  - **MongoDB** - Media uploads, folders, and content metadata
  - **PostgreSQL** - User accounts and authentication data
- **JWT Authentication** - Secure user sessions
- **LinkPreview.net Integration** - Automatic thumbnail generation

### 📊 Data Flow
1. Users authenticate and access their personal vault
2. Add media by pasting URLs from social platforms
3. Content metadata and thumbnails automatically generated
4. Organize content using folders and tags
5. Browse and search through visual grid interface

---

## 🛠️ Development Setup

### Prerequisites
- **Node.js** 18+ and npm/yarn
- **Go** 1.21+ 
- **MongoDB** (local or cloud instance)
- **PostgreSQL** (local or cloud instance)
- **Expo CLI** for mobile development

### 🏃‍♂️ Quick Start

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/lyked.git
   cd lyked
   ```

2. **Backend Setup**
   ```bash
   cd backend/
   
   # Install Go dependencies
   go mod tidy
   
   # Create environment file
   cp .env.example .env
   # Edit .env with your database credentials
   
   # Run the server
   go run cmd/main.go
   ```

3. **Frontend Setup**
   ```bash
   cd mobile/
   
   # Install dependencies
   npm install
   
   # Create environment file
   cp .env.example .env
   # Edit .env with your API endpoints
   
   # Start Expo development server
   npm start
   ```

4. **Database Setup**
   ```bash
   # MongoDB - Create database and collections
   # PostgreSQL - Run migrations (when implemented)
   ```

### 📱 Running the App

```bash
# Start backend server (Terminal 1)
cd backend && go run cmd/main.go

# Start mobile app (Terminal 2)
cd mobile && npm start

# Choose platform:
# - Press 'i' for iOS Simulator
# - Press 'a' for Android Emulator  
# - Scan QR code with Expo Go app
```

---

## 🔧 Environment Variables

### Backend (`.env`)
```bash
# Server Configuration
PORT=3000
GIN_MODE=development

# Database URLs
MONGODB_URI=mongodb://localhost:27017/lyked
POSTGRES_URI=postgres://user:password@localhost:5432/lyked

# External APIs
LINKPREVIEW_API_KEY=your_api_key_here

# Authentication (when implemented)
JWT_SECRET_KEY=your-super-secret-jwt-key
```

### Frontend (`.env`)
```bash
# API Configuration
EXPO_PUBLIC_API_BASE_URL=http://localhost:3000

# External APIs
EXPO_PUBLIC_LINKPREVIEW_API_KEY=your_api_key_here

# Authentication (when implemented)
EXPO_PUBLIC_CLERK_PUBLISHABLE_KEY=pk_test_...
```

---

## 📋 API Documentation

### Current Endpoints

#### Uploads
- `POST /uploads/upload` - Create new media item
- `GET /uploads/all?user_id=<uuid>` - Fetch user's uploads
- `DELETE /uploads/delete?id=<objectid>` - Remove upload

### Planned Endpoints
- Authentication: `/auth/*`
- Folders: `/folders/*`
- Search: `/search`
- Users: `/users/*`

*Full API documentation coming soon with Swagger/OpenAPI*

---

## 🧪 Testing

```bash
# Backend tests
cd backend && go test ./...

# Frontend tests  
cd mobile && npm test

# Linting
cd mobile && npm run lint
cd mobile && npm run format
```

---

## 📂 Project Structure

```
lyked/
├── 📱 mobile/              # React Native + Expo frontend
│   ├── src/app/           # App components and screens
│   ├── src/components/    # Reusable UI components
│   ├── src/context/       # React Context providers
│   └── src/utils/         # Helper functions
│
├── 🚀 backend/            # Go + Gin REST API
│   ├── cmd/               # Application entry points
│   ├── handlers/          # HTTP request handlers
│   ├── models/            # Database models
│   ├── routes/            # Route definitions
│   └── utils/             # Utility functions
│
├── 📋 docs/               # Documentation
└── 🔧 config/             # Configuration files
```

---

## 🤝 Contributing

We welcome contributions! Here's how to get started:

1. **Fork the repository**
2. **Create a feature branch** (`git checkout -b feature/amazing-feature`)
3. **Make your changes** following our coding conventions
4. **Test your changes** thoroughly
5. **Commit your changes** (`git commit -m 'Add amazing feature'`)
6. **Push to the branch** (`git push origin feature/amazing-feature`)
7. **Open a Pull Request**

### Development Guidelines
- Follow existing code style and conventions
- Add tests for new features
- Update documentation as needed
- Use descriptive commit messages
- Keep PRs focused and atomic

---

## 🗺️ Roadmap

Check out our detailed development roadmaps:
- [📱 Frontend Roadmap](mobile/FRONTEND_ROADMAP.md) - React Native app development
- [🚀 Backend Roadmap](backend/BACKEND_ROADMAP.md) - Go API server development

### Milestones
- **v0.1** ✅ MVP with basic media saving and viewing
- **v0.2** 🔄 User authentication and enhanced UI
- **v0.3** ⏳ Advanced search and folder management
- **v0.4** ⏳ AI recommendations and web extension
- **v1.0** ⏳ Full feature set with social capabilities

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 👥 Authors

- **Your Name** - *Initial work* - [YourGitHub](https://github.com/yourusername)

---

## 🙏 Acknowledgments

- **LinkPreview.net** - Thumbnail generation service
- **Expo Team** - Amazing React Native development platform
- **Gin Framework** - Fast and lightweight Go web framework
- **MongoDB** - Flexible document database
- **PostgreSQL** - Reliable relational database

---

## 📧 Contact & Support

- **GitHub Issues** - [Report bugs or request features](https://github.com/yourusername/lyked/issues)
- **Email** - your.email@example.com
- **Discord** - Coming soon for community discussions

---

## 🌟 Star History

[![Star History Chart](https://api.star-history.com/svg?repos=yourusername/lyked&type=Date)](https://star-history.com/#yourusername/lyked&Date)

---

<div align="center">
  <p><strong>Built with ❤️ for content curators everywhere</strong></p>
  <p>
    <a href="#-lyked">Back to Top</a> •
    <a href="https://github.com/yourusername/lyked/issues">Report Bug</a> •
    <a href="https://github.com/yourusername/lyked/issues">Request Feature</a>
  </p>
</div>
