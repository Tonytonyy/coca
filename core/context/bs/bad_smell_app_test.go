package bs

import (
	. "github.com/onsi/gomega"
	"path/filepath"
	"testing"
)

func TestBadSmellApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)

	bsApp := NewBadSmellApp()
	codePath := "../../../_fixtures/bs"
	codePath = filepath.FromSlash(codePath)

	bs := bsApp.AnalysisPath(codePath)
	bsList := bsApp.FilterBadSmell(bs, nil)

	g.Expect(len(bsList)).To(Equal(4))
	g.Expect(bsList[0].Bs).To(Equal("complexCondition"))
}
