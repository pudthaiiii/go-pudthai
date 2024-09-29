package jobs

import (
	"github.com/gocraft/work"
)

// Context สำหรับ jobs
type Context struct{}

// NewJobRegister สร้างตัวจัดการ jobs ใหม่และลงทะเบียน job ทั้งหมด
func NewJobRegister(pool *work.WorkerPool) {
	jobs := map[string]func(*work.Job) error{
		"test_job": TestJob,
	}

	// ลงทะเบียน jobs ทั้งหมด
	for name, handler := range jobs {
		pool.Job(name, handler)
	}
}
