package notes

import (
	"context"
	"errors"

	"github.com/jomei/notionapi"
)

type HttpRequest struct {
	Client *notionapi.Client
}

func (n *HttpRequest) CreateNote(title, content, databaseID string) (page *notionapi.Page, err error) {
	page, err = n.Client.Page.Create(context.Background(), &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			DatabaseID: notionapi.DatabaseID(databaseID),
			Type:       notionapi.ParentTypeDatabaseID,
		},
		Properties: notionapi.Properties{
			"Name": notionapi.PageTitleProperty{
				Title: notionapi.Paragraph{{Text: notionapi.Text{Content: title}}},
			},
		},
		Children: []notionapi.Block{
			&notionapi.ParagraphBlock{
				Type:        "paragraph",
				Object:      notionapi.ObjectTypeBlock,
				HasChildren: false,
				Paragraph: struct {
					Text     notionapi.Paragraph `json:"text"`
					Children []notionapi.Block   `json:"children,omitempty"`
				}{
					Text: notionapi.Paragraph{{Text: notionapi.Text{Content: content}}},
				},
			},
		},
	})
	return
}

func (n *HttpRequest) CreateNoteWithTags(title, content string, tags []string, databaseID string) (page *notionapi.Page, err error) {
	tagsObject := make([]notionapi.Option, 0)
	for _, tag := range tags {
		if tag != "" {
			opt := notionapi.Option{
				Name: tag,
			}
			tagsObject = append(tagsObject, opt)
		}
	}

	page, err = n.Client.Page.Create(context.Background(), &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			DatabaseID: notionapi.DatabaseID(databaseID),
			Type:       notionapi.ParentTypeDatabaseID,
		},
		Properties: notionapi.Properties{
			"Name": notionapi.PageTitleProperty{
				Title: notionapi.Paragraph{{Text: notionapi.Text{Content: title}}},
			},
			"Tags": notionapi.MultiSelectOptionsProperty{
				Type:        "array",
				MultiSelect: tagsObject,
			},
		},
		Children: []notionapi.Block{
			&notionapi.ParagraphBlock{
				Type:        "paragraph",
				Object:      notionapi.ObjectTypeBlock,
				HasChildren: false,
				Paragraph: struct {
					Text     notionapi.Paragraph `json:"text"`
					Children []notionapi.Block   `json:"children,omitempty"`
				}{
					Text: notionapi.Paragraph{{Text: notionapi.Text{Content: content}}},
				},
			},
		},
	})
	return
}

func (n *HttpRequest) GetLatestNote(databaseID string) (page *notionapi.DatabaseQueryResponse, err error) {
	page, err = n.Client.Database.Query(context.Background(), notionapi.DatabaseID(databaseID), &notionapi.DatabaseQueryRequest{
		PageSize: 1,
	})
	return
}

func (n *HttpRequest) AppendNote(pageID, content string) (page notionapi.Block, err error) {
	page, err = n.Client.Block.AppendChildren(context.Background(), notionapi.BlockID(pageID), &notionapi.AppendBlockChildrenRequest{
		Children: []notionapi.Block{
			&notionapi.ParagraphBlock{
				Type:        "paragraph",
				Object:      notionapi.ObjectTypeBlock,
				HasChildren: false,
				Paragraph: struct {
					Text     notionapi.Paragraph `json:"text"`
					Children []notionapi.Block   `json:"children,omitempty"`
				}{
					Text: notionapi.Paragraph{{Text: notionapi.Text{Content: content}}},
				},
			},
		},
	})
	return
}

func (n *HttpRequest) GetNote(databaseID, query string) (page *notionapi.DatabaseQueryResponse, err error) {
	page, err = n.Client.Database.Query(context.Background(), notionapi.DatabaseID(databaseID), &notionapi.DatabaseQueryRequest{
		PageSize: 1,
		PropertyFilter: &notionapi.PropertyFilter{
			Property: "Name",
			Text: &notionapi.TextFilterCondition{
				Equals: query,
			},
		},
	})
	return
}

func (n *HttpRequest) SearchNote(databaseID, query string) (page *notionapi.DatabaseQueryResponse, err error) {
	page, err = n.Client.Database.Query(context.Background(), notionapi.DatabaseID(databaseID), &notionapi.DatabaseQueryRequest{
		PageSize: 1,
		PropertyFilter: &notionapi.PropertyFilter{
			Property: "Name",
			Text: &notionapi.TextFilterCondition{
				Contains: query,
			},
		},
	})
	return
}

func (n *HttpRequest) FindNote(databaseID, query string) (page *notionapi.DatabaseQueryResponse, err error) {
	page, err = n.GetNote(databaseID, query)
	if len(page.Results) == 0 {
		page, err = n.SearchNote(databaseID, query)
	}
	if len(page.Results) == 0 {
		err = errors.New("Cant't find note")
	}
	return
}
