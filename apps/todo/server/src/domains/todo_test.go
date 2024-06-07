//go:build test

package domains_test

import (
	"libs/errors"
	"time"
	"todo-server/domains"

	"github.com/brianvoe/gofakeit/v6"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewTodo", func() {
	var id string
	var name string
	var desc string
	var dueAt time.Time
	var status domains.TodoStatus

	Context("create todo", func() {
		BeforeEach(func() {
			id = gofakeit.UUID()
			name = gofakeit.StreetName()
			desc = gofakeit.StreetName()
			dueAt = time.Now().AddDate(0, 0, 1)
			status = domains.TodoNotStartedStatus
		})

		It("returns todo", func() {
			todo, err := domains.NewTodo(
				id,
				name,
				desc,
				dueAt,
				status,
			)

			Expect(err).NotTo(HaveOccurred())
			Expect(todo.ID()).To(Equal(id))
			Expect(todo.Name()).To(Equal(name))
			Expect(todo.Description()).To(Equal(desc))
			Expect(todo.DueAt()).To(Equal(dueAt))
			Expect(todo.Status()).To(Equal(status))
		})

		It("throws validation error with invalid name", func() {
			name = "op"

			_, err := domains.NewTodo(
				id,
				name,
				desc,
				dueAt,
				status,
			)

			Expect(err).To(HaveOccurred())
			Expect(err.(errors.Error).Code()).To(Equal(errors.ErrorEntityValidation))
		})

		It("throws validation error with invalid dueAt", func() {
			dueAt = time.Now().AddDate(0, 0, -2)

			_, err := domains.NewTodo(
				id,
				name,
				desc,
				dueAt,
				status,
			)

			Expect(err).To(HaveOccurred())
			Expect(err.(errors.Error).Code()).To(Equal(errors.ErrorEntityValidation))
		})

	})
})
