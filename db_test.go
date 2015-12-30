package hyperdb_test

import (
	"github.com/mysza/hyperdb"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os"
)

var _ = Describe("DB", func() {
	Describe("Opening the database", func() {
		Context("with defaults", func() {
			It("should open default database", func() {
				// calling Open with empty string should use default file name
				db, err := hyperdb.Open("")
				_, pathErr := os.Stat(hyperdb.DefaultDataFileName)
				defer func() {
					db.Close()
					os.Remove(hyperdb.DefaultDataFileName)
				}()
				Expect(db).ToNot(BeNil())
				Expect(err).To(BeNil())
				Expect(pathErr).To(BeNil())
			})
			It("should open named database", func() {
				const dbname = "very_unique_name.db"
				db, err := hyperdb.Open(dbname)
				_, pathErr := os.Stat(dbname)
				defer func() {
					db.Close()
					os.Remove(dbname)
				}()
				Expect(db).ToNot(BeNil())
				Expect(err).To(BeNil())
				Expect(pathErr).To(BeNil())
			})
		})
	})
})
