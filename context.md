# 💡 Project Title: Lyked

**Lyked** is a mobile application that empowers users to save and organize their favorite videos and images from various social media platforms — such as TikTok, Instagram, YouTube, and Pinterest — in a private, personalized media vault. The app gives users a clean, intuitive experience to curate content they resonate with and revisit it whenever they want.

---

## 🧩 Problem

In the age of infinite scrolling, users often come across inspiring, funny, emotional, or educational content — only to lose it moments later. Social media apps aren’t designed for long-term media curation. Users resort to scattered bookmarks, screenshots, or saved folders that are platform-locked and difficult to manage.

---

## ✅ Solution

**Lyked** gives users one unified space to:

- Collect links to favorite images and videos from across the internet.
- Categorize and tag media for easy filtering.
- View all saved content in a scrollable visual library.
- Organize content without relying on the original platforms' saved-post features.

> _Future versions will include onboard AI that learns from stored media and recommends similar videos and images from the internet — tailoring discovery to the user's true tastes._

---

## 🚀 Key Features (MVP)

- **User Authentication** via Clerk (email & social login)
- **Add Media**: Paste a link to an image or video, give it a title and optional tags
- **Categorized Media Feed**: Browse saved content by type (e.g., memes, style, inspiration)
- **Delete or Edit Entries**: Maintain a curated library
- **Mobile-first UI**: Built using React Native + Expo

---

## 🏗️ Tech Stack

### Frontend (Mobile App)

- React Native with Expo
- Clerk for authentication
- React Context API for state management
- Axios for API requests

### Backend

- Golang with Gin framework
- JWT for secure user sessions
- PostgreSQL via GORM for user data storage
- MongoDB for folders and uploads
- RESTful API design
- Cloudinary or AWS S3 for media previews (if extended)

### Infrastructure

- Railway / Render for backend hosting
- Neon / Supabase for PostgreSQL database
- Monorepo with Yarn Workspaces or Turborepo

---

## 🛣️ Future Roadmap

- ✅ Save, view, and organize media
- ⏳ Full-text search and filtering by tags
- ⏳ Image preview thumbnails
- ⏳ AI-powered recommendations (via vector embeddings)
- ⏳ Web browser extension for 1-click saving
- ⏳ Social features (share collections with friends)

---
