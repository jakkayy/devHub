# devHub Infrastructure Setup

ไฟล์ `docker-compose.yml` นี้ใช้สำหรับ Spin up **Infrastructure Services (PostgreSQL & Redis)** สำหรับการพัฒนาในเครื่อง (Local Development)

## 🚀 วิธีการรัน Infrastructure

1. สั่งรัน PostgreSQL และ Redis:
   ```bash
   cd deploy
   docker compose up -d
   ```

2. ตรวจสอบการรัน Container:
   ```bash
   docker compose ps
   ```

3. วิธีรัน แอปพลิเคชันในเครื่อง (Native Development):
   * **Go Backend Service:**
     ```bash
     cd services/backend-api
     go run cmd/api/main.go
     ```
   * **Next.js Web Frontend:**
     ```bash
     cd apps/web
     bun run dev
     ```

4. สั่งหยุดการทำงานของ Infrastructure:
   ```bash
   cd deploy
   docker compose down
   ```
