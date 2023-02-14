package entity

import (
	"testing"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
	"gorm.io/gorm"
)
type Supaporn struct{
	gorm.Model
	Name string    `valid:"required~Name cannot be blank"`
}

func TestDataVideoValidate(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check data is valid", func(t *testing.T) {
		u := Supaporn{
			Name: "Supaporn Seangduean",
		}
		ok, err := govalidator.ValidateStruct(u)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}
func TestNameVideoValidate(t *testing.T) {
		g := NewGomegaWithT(t)
	t.Run("check name not blank", func(t *testing.T) {
		u := Supaporn{
			Name: "", //ผิด
		}
		ok, err := govalidator.ValidateStruct(u)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Name cannot be blank"))
	})
}