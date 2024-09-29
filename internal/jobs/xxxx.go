package jobs

import (
	"log"

	"github.com/gocraft/work"
)

// TestJob ตัวอย่างสำหรับการประมวลผลงาน
func TestJob(job *work.Job) error {
	name := job.ArgString("name")
	log.Printf("Processing TestJob with name: %s", name)
	return nil
}
