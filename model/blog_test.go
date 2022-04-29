package model_test

import (
	"testing"
	"time"
	
	"github.com/pandudpn/blog/model"
	"github.com/stretchr/testify/assert"
)

func TestBlog(t *testing.T) {
	testCases := []struct {
		name                 string
		blog                 *model.Blog
		expectedResultStatus string
		expectedIsPublish    bool
		expectedIsDraft      bool
	}{
		{
			name: "Test Case #1 Blog Published",
			blog: &model.Blog{
				Id:        1,
				Title:     "Blog unit test",
				Body:      "This is used for unit testing",
				Status:    1,
				CreatedBy: 1,
				CreatedAt: time.Now(),
				Image:     "https://image.com",
			},
			expectedResultStatus: "Published",
			expectedIsDraft:      false,
			expectedIsPublish:    true,
		},
		{
			name: "Test Case #1 Blog Published",
			blog: &model.Blog{
				Id:        1,
				Title:     "Blog unit test",
				Body:      "This is used for unit testing",
				Status:    0,
				CreatedBy: 1,
				CreatedAt: time.Now(),
				Image:     "https://image.com",
			},
			expectedResultStatus: "Draft",
			expectedIsDraft:      true,
			expectedIsPublish:    false,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status := tc.blog.GetStatus()
			isPublished := tc.blog.IsStatusPublish()
			isDrafted := tc.blog.IsStatusDraft()
			
			assert.Equal(t, tc.expectedResultStatus, status)
			assert.Equal(t, tc.expectedIsPublish, isPublished)
			assert.Equal(t, tc.expectedIsDraft, isDrafted)
		})
	}
}
