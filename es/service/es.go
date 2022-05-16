package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/es/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/es/model/response"
	"github.com/olivere/elastic/v7"
)

type ElasticsearchService struct{}

//
// CreateElasticsearch
//  @Description: 新增Elasticsearch数据
//  @param name 索引名
//  @param esId 索引ID
//  @param data 数据
//  @return err 错误
//  @return ret 返回值
//
func (e *ElasticsearchService) CreateElasticsearch(name string, esId string, data interface{}) (err error, ret interface{}) {

	get, err := global.Elasticsearch.Index().Index(name).Id(esId).BodyJson(data).Do(context.Background())
	if err != nil {
		return err, nil
	}
	return err, get
}

//
// DeleteElasticsearch
//  @Description: 删除Elasticsearch数据
//  @param name 索引名
//  @param esId 索引ID
//  @return err
//  @return ret
//
func (e *ElasticsearchService) DeleteElasticsearch(name string, esId string) (err error, ret interface{}) {

	get, err := global.Elasticsearch.Delete().Index(name).Id(esId).Do(context.Background())
	if err != nil {
		return err, nil
	}
	return err, get
}

//
// UpdateElasticsearch
//  @Description:
//  @param name 索引名
//  @param esId 索引ID
//  @param data 数据
//  @return err 错误
//  @return ret 返回值
//
func (e *ElasticsearchService) UpdateElasticsearch(name string, esId string, data interface{}) (err error, ret interface{}) {

	get, err := global.Elasticsearch.Update().Index(name).Id(esId).Doc(data).Do(context.Background())
	if err != nil {
		return err, nil
	}
	return err, get
}

//
// GetIdElasticsearch
//  @Description: 通过id查找Elasticsearch
//  @param name 索引名
//  @param esId 索引ID
//  @return err 错误
//  @return ret 返回值
//
func (e *ElasticsearchService) GetIdElasticsearch(name string, esId string) (err error, ret interface{}) {
	get, err := global.Elasticsearch.Get().Index(name).Id(esId).Do(context.Background())
	if err != nil {
		return err, nil
	}
	return err, get
}

//
// GetCountElasticsearch
//  @Description: 当前ID索引数量
//  @param name 索引名
//  @return ret 返回值
//
func (e *ElasticsearchService) GetCountElasticsearch(name string) (ret interface{}) {
	list, err := global.Elasticsearch.Count(name).Do(context.Background())
	if err != nil {
		return err
	}
	return list
}

//
// GetQueryElasticsearch
//  @Description: 查询Elasticsearch
//  @param info 查询条件
//  @param name 索引名
//  @param text 查询字段
//  @return err 错误
//  @return ret 返回值
//
func (e *ElasticsearchService) GetQueryElasticsearch(info response.ElasticSearchSearch, name, searchField string) (err error, ret interface{}) {
	size := info.PageSize
	page := info.Page
	//根据name索引查询Elasticsearch数据
	boolQ := elastic.NewQueryStringQuery(info.Title)
	boolQ = boolQ.Field(searchField)
	get, err := global.Elasticsearch.Search(name).
		Query(boolQ). // specify the query
		//Sort("id", true). //按字段"age"排序，升序排列
		Size(size). // 分页，单页显示10条
		From((page - 1) * size).
		//FetchSourceContext(fsc).//只取对应字段
		Do(context.Background()) // 执行
	if err != nil {
		return err, ""
	}
	return err, get
}
