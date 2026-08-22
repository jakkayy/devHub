---
name: clean-architecture-checker
description: Guideline and rules for enforcing Clean Architecture, Modular Layering, and Performance Optimization across Go Backend, Next.js Frontend, and VS Code Extension in devHub project.
---

# Clean Architecture & Code Quality Skill

คำแนะนำและขั้นตอนการตรวจสอบโครงสร้างโค้ดและ Architecture ของโปรเจกต์ `devHub` เพื่อให้เป็นระเบียบ อ่านง่าย ขยายง่าย และทรงประสิทธิภาพ

## 🎯 หลักการสถาปัตยกรรม (Architecture Checklist)

### 1. Go Backend (`services/backend-api`)
- **Domain Layer (`internal/domain`)**: ต้องเป็น Pure Go Structs และ Interfaces ห้าม Import External Frameworks (เช่น Gin, GORM)
- **Usecase / Service Layer (`internal/usecase`)**: ควบคุม Business Logic และรับ-ส่งข้อมูลผ่าน Interfaces เท่านั้น
- **Repository Layer (`internal/repository`)**: ติดต่อกับ Database (PostgreSQL/Redis) หรือ External APIs
- **Handler / Delivery Layer (`internal/handler`)**: รับ HTTP Requests (Gin/Fiber), Validate DTOs และส่งข้อมูลกลับเป็น JSON
- **Concurrency Safety**: 
  - การดึงข้อมูลขนานกันจาก 6 APIs ต้องใช้ `errgroup.Group` หรือ `sync.WaitGroup` พร้อมกับการตั้ง `context.WithTimeout` (ไม่เกิน 5-10 วินาที)

### 2. Next.js Frontend (`apps/web`)
- **Single Responsibility Components**: 1 Component ควรทำหน้าที่เดียว ไม่ควรมีไฟล์ยาวเกิน 250-300 บรรทัด
- **Custom Hooks Isolation**: แยก Logic Data Fetching ออกไปเป็น Custom Hooks (เช่น `useGitHubIssues()`, `useGoogleSheetTasks()`)
- **Type Safety**: ใช้ TypeScript Interfaces/Types ที่ Sync ตรงกันกับ Go DTOs

### 3. VS Code Extension (`apps/extension`)
- **Lightweight Touchpoint**: ทำหน้าที่เป็นเพียง Presentation & Event Handling แล้วสั่งคำสั่งไปยัง Go Backend
- **URI Protocol Handler**: รองรับ `vscode://file/path:line` สำหรับการทำ Deep Linking จาก Web Dashboard

---

## ⚡ คำแนะนำสำหรับการ Optimize (Performance Best Practices)

1. **Redis Caching Strategy**:
   - ข้อมูลจาก External APIs (เช่น Miro Diagram, Sheets, API Specs) ควรทำ Cache ไว้ใน Redis โดยมี TTL ที่เหมาะสม (เช่น 1-5 นาที) เพื่อป้องกัน API Rate Limit
2. **Database Indexing**:
   - ตารางที่ใช้เชื่อมโยง (Link Mapping: Task ID <-> Git Branch <-> API Endpoint) ต้องทำ Compound Index ไว้เสมอ
3. **Frontend Lazy Loading**:
   - ส่วนประกอบหนักๆ เช่น Miro iFrame Embed หรือ Interactive API Explorer ให้ใช้ `next/dynamic` ทำ Lazy Loading เสมอ
