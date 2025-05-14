# Armada Management System

Sistem ini memantau lokasi kendaraan secara real-time melalui MQTT, menyimpan data ke PostgreSQL, REST API, dan mengirim event masuk geofence ke RabbitMQ.

## 🧱 Arsitektur

- **MQTT Broker**: eclipse-mosquitto
- **Database**: PostgreSQL
- **Message Broker**: RabbitMQ
- **Backend**: Go
- **Container Orchestration**: Docker Compose

## 🚀 Cara Menjalankan

### 1. Clone Repository

```bash
git clone <repository-url>
cd <repository-folder>
```

### 2. Tambahkan env dan jalankan Docker Compose

Buat file .env lalu copy paste yang ada diexample

Jika tidak ada yang ingin dirubah bisa jalankan docker compose:

```bash
docker compose up --build
```

Tunggu hingga semua container berjalan tanpa error. Akses:

- RabbitMQ dashboard: [http://localhost:15672](http://localhost:15672)

  - Username: `armada_user`
  - Password: `armada_pass`

- API (Backend): [http://localhost:8080](http://localhost:8080)

### 3. Jalankan Mock Publisher

```bash
docker compose exec armada-app sh
go run mock/publisher.go
```

File `mock/publisher.go` akan mengirimkan data lokasi kendaraan ke MQTT broker secara berkala.

---

## 🔍 Pengujian & Hasil yang Diharapkan

| Fitur                       | Deskripsi                                                          | Hasil yang Diharapkan                                           |
| --------------------------- | ------------------------------------------------------------------ | --------------------------------------------------------------- |
| **Publikasi data MQTT**     | `mock/publisher.go` mengirim data lokasi kendaraan ke broker       | Data dikirim dan diterima oleh backend                          |
| **Penyimpanan PostgreSQL**  | Backend menyimpan data ke tabel `vehicle_locations`                | Data lokasi tersimpan di database secara akurat                 |
| **API Lokasi Terakhir**     | Endpoint: `GET /vehicles/{vehicle_id}/location`                    | Mengembalikan data lokasi terbaru kendaraan                     |
| **API Riwayat Lokasi**      | Endpoint: `GET /vehicles/{vehicle_id}/history?start=unix&end=unix` | Mengembalikan riwayat lokasi dalam rentang waktu tertentu       |
| **RabbitMQ Geofence Event** | Backend kirim event ke RabbitMQ saat kendaraan masuk area geofence | Pesan berhasil dikirim dan bisa dilihat dari RabbitMQ dashboard |
| **Docker Compose**          | Seluruh layanan berjalan dengan konfigurasi `docker-compose.yml`   | Tidak ada container yang error atau restart terus-menerus       |

---

## 📬 Endpoint API

- **Lokasi terakhir kendaraan**
  GET /vehicles/{vehicle_id}/location

- **Riwayat lokasi kendaraan**
  GET /vehicles/{vehicle_id}/history?start=<unix>&end=<unix>

Contoh:

```bash
curl http://localhost:8080/vehicles/B1234XYZ/location
curl http://localhost:8080/vehicles/B1234XYZ/history?start=1747200000&end=1747209999
```

Berikut untuk Postman Collection:
https://drive.google.com/file/d/11wKY413wWnRv7QVnbPR0kTYKUOSuqBMB/view?usp=sharing

---

## 🧹 Shutdown

Untuk menghentikan semua layanan:

```bash
docker compose down
```
