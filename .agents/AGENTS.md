# Project Customization & Coding Guidelines (`AGENTS.md`)

## 🏛️ 1. สถาปัตยกรรมและโครงสร้างโปรเจกต์ (Clean Architecture Guidelines)

โปรเจกต์ **DevHub (Internal Developer Portal)** ถูกออกแบบด้วยสถาปัตยกรรม **Modular Clean Architecture** รองรับการปรับแต่ง (Optimization) และขยายระบบ (Scalability) ในอนาคต

### 📁 1.1 โครงสร้างโฟลเดอร์หลัก (Monorepo / Multi-service Pattern)
```
devHub/
├── .agents/                      # Custom Rules & Skills สำหรับ AI coding assistants
│   ├── AGENTS.md                 # กฎและ Coding Standards (ไฟล์นี้)
│   └── skills/                   # Custom Project Skills
├── apps/                         # Applications & Touchpoints
│   ├── web/                      # Next.js 14 (App Router) + Tailwind + Shadcn UI
│   └── extension/                # VS Code Extension (TypeScript)
├── services/                     # Backend Services
│   └── backend-api/              # Go (Golang) Microservice (Gin/Fiber + GORM)
├── deploy/                       # Infrastructure & Deployment
│   ├── docker-compose.yml        # Local development setup (Go, Next.js, Postgres, Redis)
│   └── docker/                   # Dockerfiles (Multi-stage builds)
└── README.md                     # Project documentation & Overview
```

## 🌿 1.2 กฎการตั้งชื่อ Branch (Git Branching Model)

โปรเจกต์นี้ใช้โครงสร้าง **Trunk-Based / Feature Branching Strategy** โดยมีการกำหนดรูปแบบการตั้งชื่อ Branch ที่ชัดเจนดังนี้:

* **`main`**: Production-ready code (เสถียรที่สุด)
* **`develop`**: Integration branch หลักสำหรับรันพัฒนาประจำวัน
* **`feature/<component>-<short-description>`**: สำหรับฟีเจอร์ใหม่
  * ตัวอย่าง: `feature/backend-db-migration`, `feature/web-dashboard-ui`, `feature/extension-sidebar`
* **`fix/<component>-<short-description>`**: สำหรับแก้บั๊ก
  * ตัวอย่าง: `fix/backend-redis-timeout`, `fix/web-layout-overflow`
* **`ci/<description>`**: สำหรับงาน CI/CD & DevOps Pipeline
  * ตัวอย่าง: `ci/github-actions-pipeline`
* **`docs/<description>`**: สำหรับงานปรับแต่ง Documentation
  * ตัวอย่าง: `docs/update-api-spec`


ใช้แนวคิด **Clean Architecture & Idiomatic Go**:
* **Layer Isolation:**
  * `domain/` / `entity/`: Data Models และ Interfaces หลัก (ไม่มี external dependencies)
  * `usecase/` / `service/`: Business Logic หลักของระบบ
  * `repository/`: การติดต่อกับ Database (GORM/PostgreSQL, Redis)
  * `delivery/` / `handler/`: REST API Handlers (Gin/Fiber) & DTOs
* **Concurrency:** 
  * ใช้ Goroutines และ Worker Pools ในการทำ Parallel Fetching ข้อมูลจาก External APIs (GitHub, Sheets, Discord, Miro)
  * จัดการ Context Timeout (`context.Context`) เสมอเมื่อเรียก External APIs
* **Error Handling:** 
  * ห้าม Swallow Errors สุ่มสี่สุ่มห้า ต้องใช้วิธี Explicit Error Wrapping (`fmt.Errorf("...: %w", err)`)
  * ป้องกัน Panic ด้วย `recover()` middleware ใน Handler layer

---

## ⚛️ 3. กฎการเขียนโค้ดสำหรับ Next.js Frontend (`apps/web`)

เน้น **Component Driven & Performance First**:
* **Directory Structure:**
  * `app/`: Next.js App Router Pages & Layouts
  * `components/ui/`: Reusable primitives (Shadcn UI)
  * `components/features/`: Component เฉพาะตาม Feature (e.g., `github/`, `sheets/`, `api-spec/`)
  * `services/` / `hooks/`: API Clients & Custom React Query Hooks (`TanStack Query`)
* **State Management & Data Fetching:**
  * ใช้ Server Components (RSC) ในส่วนที่เป็น Static/SEO Data
  * ใช้ Client Components (`"use client"`) เฉพาะจุดที่มีการโต้ตอบ (Interactive UI)
  * ใช้ TanStack Query สำหรับ Caching & Auto-refreshing 6 External Tools
* **Design & Styling:**
  * คุมโทน **Developer-Centric Dark Mode** พรีเมียม (คล้าย Vercel / Linear / Raycast)
  * ห้าม Hardcode Inline Styles หรือ Static Pixel Heights; ให้ใช้ Tailwind Tokens

---

## ⚡ 3.5 Package Manager Policy (Bun Only)

โปรเจกต์นี้กำหนดให้ใช้ **Bun** (`bun`) เป็น Package Manager และ Runtime หลักสำหรับ Node.js/TypeScript Ecosystem ทั้งหมด (ทั้งใน `apps/web` และ `apps/extension`):
* **คำสั่งติดตั้ง:** ใช้ `bun install` แทน `npm install`
* **คำสั่งรัน script:** ใช้ `bun run <script>` แทน `npm run`
* **คำสั่งรัน package:** ใช้ `bunx <package>` แทน `npx`


* **Performance Isolation:** 
  * ห้ามรันงานซับซ้อน หรือ Synchronous Blocking บน Extension Host Main Thread
  * งานประมวลผลขนาดใหญ่ให้ส่งกลับไปทำที่ **Go Backend Service** เสมอ
* **Webview Communication:** 
  * ใช้ `postMessage` และ Strict Type Messaging ในการสื่อสารระหว่าง Extension Host กับ Webview UI

---

## 🐳 5. กฎสถาปัตยกรรม Docker & Optimization

* **Multi-stage Builds:** ต้องใช้ Multi-stage Build ทุก Dockerfile เพื่อให้ได้ขนาด Image เล็กที่สุด (Go < 20MB, Next.js Standalone < 150MB)
* **Resource Constraints:** กำหนด memory/cpu limits ใน `docker-compose.yml` เพื่อให้รันในเครื่อง Dev ได้เบาและเสถียร
