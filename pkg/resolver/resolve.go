package resolver

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"strconv"
	"strings"
)

var comments = make([]Comment, 0, 30)

func ResolveComments(r io.Reader) ([]Comment, error) {

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatalln(err.Error())
	}

	doc.Find(".comment-item").Each(resolveComment)

	return comments, nil
}

func resolveComment(_ int, selection *goquery.Selection) {
	cid, err := resolveCid(selection)
	if err != nil {
		log.Panicln(err.Error())
	}

	log.Printf("[ Resolving ] Resolving Comment #%s", cid)

	authorId, err := resolveAuthorId(selection)
	if err != nil {
		log.Panicln(err.Error())
	}

	authorName, err := resolveAuthorName(selection)
	if err != nil {
		log.Panicln(err.Error())
	}

	authorAvatar, err := resolveAuthorAvatar(selection)
	if err != nil {
		log.Panicln(err.Error())
	}

	rate, err := resolveRate(selection)
	if err != nil {
		log.Panicln(err.Error())
	}

	date, err := resolveDate(selection)
	if err != nil {
		log.Panicln(err.Error())
	}

	content, err := resolveContent(selection)
	if err != nil {
		log.Panicln(err.Error())
	}

	vote, err := resolveVote(selection)
	if err != nil {
		log.Panicln(err.Error())
	}

	comment := Comment{
		Cid:          cid,
		AuthorId:     authorId,
		AuthorName:   authorName,
		AuthorAvatar: authorAvatar,
		Rate:         rate,
		Date:         date,
		Content:      content,
		Vote:         vote,
	}

	comments = append(comments, comment)
}

func resolveCid(selection *goquery.Selection) (string, error) {
	v, e := selection.Attr("data-cid")
	if e != true {
		return "", errors.New("cid not found")
	} else {
		return v, nil
	}
}

func resolveAuthorId(selection *goquery.Selection) (string, error) {
	v, e := selection.Find(".comment-info a").Attr("href")
	if e != true {
		return "", errors.New("author id not found")
	} else {
		s := strings.Split(v, "/")
		return s[len(s)-1], nil
	}
}

func resolveAuthorName(selection *goquery.Selection) (string, error) {
	v := selection.Find(".comment-info a").Text()
	if len(v) == 0 {
		return "", errors.New("author name not found")
	} else {
		return v, nil
	}
}

func resolveAuthorAvatar(selection *goquery.Selection) (string, error) {
	v, e := selection.Find(".avatar img").Attr("src")
	if e != true {
		return "", errors.New("author avatar not found")
	} else {
		return v, nil
	}
}

func resolveRate(selection *goquery.Selection) (int, error) {
	v, e := selection.Find(".comment-info .rating").Attr("class")
	if e != true {
		return 0, errors.New("rating not found")
	} else {
		s := strings.Split(v, " ")
		r := s[0]
		switch r {
		case "allstar50":
			return 10, nil
		case "allstar40":
			return 8, nil
		case "allstar30":
			return 6, nil
		case "allstar20":
			return 4, nil
		case "allstar10":
			return 2, nil
		default:
			return 0, errors.New("rating not valid")
		}
	}
}

func resolveDate(selection *goquery.Selection) (string, error) {
	v, e := selection.Find(".comment-info .comment-time").Attr("title")
	if e != true {
		return "", errors.New("date not found")
	} else {
		return v, nil
	}
}

func resolveContent(selection *goquery.Selection) (string, error) {
	v := selection.Find(".comment-content .short").Text()
	if len(v) == 0 {
		return "", errors.New("content not found")
	} else {
		return v, nil
	}
}

func resolveVote(selection *goquery.Selection) (int, error) {
	v := selection.Find(".vote-count").Text()
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, errors.New("vote num not found")
	} else {
		return i, nil
	}
}
