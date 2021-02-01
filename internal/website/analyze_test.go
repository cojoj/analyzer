package website

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/cojoj/analyzer/internal/models"
	"golang.org/x/net/html"
)

func document(s string) *goquery.Document {
	return goquery.NewDocumentFromNode(node(s))
}

func node(s string) *html.Node {
	node, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}

	return node
}

func Test_pageTitle(t *testing.T) {
	type args struct {
		d *goquery.Document
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Document with title",
			args: args{document(`<!DOCTYPE html><html><body><title>Title</title><h1>My First Heading</h1></body></html>`)},
			want: "Title",
		},
		{
			name: "Document without title",
			args: args{document(`<!DOCTYPE html><html><body><h1>My First Heading</h1></body></html>`)},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pageTitle(tt.args.d); got != tt.want {
				t.Errorf("pageTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_documentType(t *testing.T) {
	type args struct {
		d *goquery.Document
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "HTML 5 Document",
			args: args{document(`<!DOCTYPE html><html></html>`)},
			want: `<!DOCTYPE html>`,
		},
		{
			name: "HTML 4 Document",
			args: args{document(`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "https://www.w3.org/TR/html4/loose.dtd">`)},
			want: `<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "https://www.w3.org/TR/html4/loose.dtd">`,
		},
		{
			name: "Missing DOCTYPE",
			args: args{document(`<html><body><h1>My First Heading</h1></body></html>`)},
			want: `Not specified`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := documentType(tt.args.d); got != tt.want {
				t.Errorf("documentType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasLoginForm(t *testing.T) {
	type args struct {
		d *goquery.Document
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Has login form",
			args: args{document(`<html><body><form action="/action_page" method="post"><input type="password">Test</input></form></body></html>`)},
			want: true,
		},
		{
			name: "Doesn't have login form",
			args: args{document(`<html><body><form action="/action_page" method="post"><input type="button">Test</input></form></body></html>`)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasLoginForm(tt.args.d); got != tt.want {
				t.Errorf("hasLoginForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headings(t *testing.T) {
	type args struct {
		d *goquery.Document
	}
	tests := []struct {
		name string
		args args
		want []models.Heading
	}{
		{
			name: "Single h1",
			args: args{document(`<html><body><h1>H1</h1></body></html>`)},
			want: []models.Heading{{Level: "h1", Amount: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := headings(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("headings() = %v, want %v", got, tt.want)
			}
		})
	}
}
