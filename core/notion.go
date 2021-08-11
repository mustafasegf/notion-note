package core

import (
	"context"

	"github.com/jomei/notionapi"
)

type Notion struct {
	Client *notionapi.Client
}

func (n *Notion) CreateNote(title, content, databaseID string) (page *notionapi.Page, err error) {
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

func (n *Notion) CreateNoteWithTags(title, content string, tags []string, databaseID string) (page *notionapi.Page, err error) {
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

func (n *Notion) GetLatestNote(databaseID string) (page *notionapi.DatabaseQueryResponse, err error) {
	page, err = n.Client.Database.Query(context.Background(), notionapi.DatabaseID(databaseID), &notionapi.DatabaseQueryRequest{
		PageSize: 1,
	})
	return
}

func (n *Notion) AppendNote(pageID, content string) (page notionapi.Block, err error) {
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
