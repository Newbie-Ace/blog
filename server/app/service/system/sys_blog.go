package system

import (
	"errors"
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/model/common/request"
	"github.com/noobHuKai/app/model/system"
	"gorm.io/gorm"
)

type BlogService struct{}

// CreateBlog 创建博客
func (service *BlogService) CreateBlog(s system.SysBlog, t []uint) (err error) {
	err = g.DB.Create(&s).Error
	if err != nil {
		return err
	}
	var tags []*system.SysBlogTag
	for _, v := range t {
		tags = append(tags, &system.SysBlogTag{TagID: v, BlogID: s.ID})
	}
	err = g.DB.Create(&tags).Error

	return err
}

func (service *BlogService) ListBlog(p request.Page) ([]system.BlogContent, error) {
	var blogs []system.SysBlog
	var listBlog []system.BlogContent
	err := g.DB.Find(&blogs).Offset(p.PageNum).Limit(p.PageSize).Error
	if err != nil {
		g.Logger.Error(err.Error())
		return nil, err
	}
	for _, v := range blogs {
		var tags []system.SysTag
		err = g.DB.Model(&system.SysTag{}).Joins("left join sys_blog_tag on sys_blog_tag.tag_id = sys_tag.id").Where("blog_id = ?", v.ID).Find(&tags).Error
		if err != nil {
			g.Logger.Error(err.Error())
			return nil, err
		}
		listBlog = append(listBlog, system.BlogContent{
			Blog: v,
			Tags: tags,
		})
	}
	return listBlog, nil
}

func (service *BlogService) BlogContent(id uint) (*system.BlogContent, error) {
	var blog system.SysBlog
	err := g.DB.First(&blog, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, g.ErrUserNotFound
		}
		return nil, err
	}

	var tags []system.SysTag

	err = g.DB.Model(&system.SysTag{}).Joins("left join sys_blog_tag on sys_blog_tag.tag_id = sys_tag.id").Where("blog_id = ?", blog.ID).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	res := &system.BlogContent{
		Blog: blog,
		Tags: tags,
	}
	return res, nil
}

// CreateBlogTag 创建博客tag
func (service *BlogService) CreateBlogTag(s system.SysTag) (system.SysTag, error) {
	err := g.DB.Create(&s).Error
	return s, err
}

// CreateBlogType 创建博客类型
func (service *BlogService) CreateBlogType(s system.SysBlogType) (system.SysBlogType, error) {
	err := g.DB.Create(&s).Error
	return s, err
}
