package repository

import (
	"errors"

	"xuetu-project/internal/model"

	"gorm.io/gorm"
)

// SubjectInfoRepository 提供题目及其关联实体的持久化操作
type SubjectInfoRepository interface {
	BeginTransaction() *gorm.DB
	Transaction(func(tx *gorm.DB) error) error

	CreateSubjectInfo(info *model.SubjectInfo, tx *gorm.DB) error
	UpdateSubjectInfo(info *model.SubjectInfo, tx *gorm.DB) error
	DeleteSubjectInfo(id uint64, tx *gorm.DB) error
	GetSubjectInfoByID(id uint64, tx *gorm.DB) (*model.SubjectInfo, error)

	CreateBrief(brief *model.Brief, tx *gorm.DB) error
	GetBriefBySubjectID(subjectID uint64, tx *gorm.DB) (*model.Brief, error)

	CreateJudge(judge *model.Judge, tx *gorm.DB) error
	GetJudgeBySubjectID(subjectID uint64, tx *gorm.DB) (*model.Judge, error)

	BatchCreateRadios(radios []*model.Radio, tx *gorm.DB) error
	ListRadiosBySubjectID(subjectID uint64, tx *gorm.DB) ([]*model.Radio, error)

	BatchCreateMultiples(multiples []*model.Multiple, tx *gorm.DB) error
	ListMultiplesBySubjectID(subjectID uint64, tx *gorm.DB) ([]*model.Multiple, error)

	CreateCategory(category *model.SubjectCategory, tx *gorm.DB) error
	GetCategoryByID(id uint64, tx *gorm.DB) (*model.SubjectCategory, error)
	ListCategoriesByParent(parentID *uint64, tx *gorm.DB) ([]*model.SubjectCategory, error)

	CreateLabel(label *model.SubjectLabel, tx *gorm.DB) error
	GetLabelByID(id uint64, tx *gorm.DB) (*model.SubjectLabel, error)
	ListLabels(tx *gorm.DB) ([]*model.SubjectLabel, error)

	CreateMapping(mapping *model.SubjectMapping, tx *gorm.DB) error
	ListMappingsBySubjectID(subjectID uint64, tx *gorm.DB) ([]*model.SubjectMapping, error)
	DeleteMappingsBySubjectID(subjectID uint64, tx *gorm.DB) error

	CreateLiked(liked *model.SubjectLiked, tx *gorm.DB) error
	GetLiked(subjectID uint64, likeUserID string, tx *gorm.DB) (*model.SubjectLiked, error)
	UpdateLikedStatus(subjectID uint64, likeUserID string, status int, tx *gorm.DB) error
}

type subjectInfoRepository struct {
	db *gorm.DB
}

// NewSubjectInfoRepository 构造题目仓库实现
func NewSubjectInfoRepository(db *gorm.DB) SubjectInfoRepository {
	return &subjectInfoRepository{db: db}
}

// BeginTransaction 开启事务
func (r *subjectInfoRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}

// Transaction 在事务中执行函数
func (r *subjectInfoRepository) Transaction(txFunc func(tx *gorm.DB) error) error {
	return r.db.Transaction(txFunc)
}

// CreateSubjectInfo 创建题目信息
func (r *subjectInfoRepository) CreateSubjectInfo(info *model.SubjectInfo, tx *gorm.DB) error {
	return r.withTx(tx).Create(info).Error
}

// UpdateSubjectInfo 更新题目信息
func (r *subjectInfoRepository) UpdateSubjectInfo(info *model.SubjectInfo, tx *gorm.DB) error {
	return r.withTx(tx).Save(info).Error
}

// DeleteSubjectInfo 删除题目信息（软删除）
func (r *subjectInfoRepository) DeleteSubjectInfo(id uint64, tx *gorm.DB) error {
	return r.withTx(tx).Delete(&model.SubjectInfo{}, id).Error
}

// GetSubjectInfoByID 根据 ID 查询题目
func (r *subjectInfoRepository) GetSubjectInfoByID(id uint64, tx *gorm.DB) (*model.SubjectInfo, error) {
	var info model.SubjectInfo
	if err := r.withTx(tx).First(&info, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &info, nil
}

// CreateBrief 创建简答题答案
func (r *subjectInfoRepository) CreateBrief(brief *model.Brief, tx *gorm.DB) error {
	return r.withTx(tx).Create(brief).Error
}

// GetBriefBySubjectID 按题目查询简答题答案
func (r *subjectInfoRepository) GetBriefBySubjectID(subjectID uint64, tx *gorm.DB) (*model.Brief, error) {
	var brief model.Brief
	if err := r.withTx(tx).Where("subject_id = ?", subjectID).First(&brief).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &brief, nil
}

// CreateJudge 创建判断题答案
func (r *subjectInfoRepository) CreateJudge(judge *model.Judge, tx *gorm.DB) error {
	return r.withTx(tx).Create(judge).Error
}

// GetJudgeBySubjectID 按题目查询判断题答案
func (r *subjectInfoRepository) GetJudgeBySubjectID(subjectID uint64, tx *gorm.DB) (*model.Judge, error) {
	var judge model.Judge
	if err := r.withTx(tx).Where("subject_id = ?", subjectID).First(&judge).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &judge, nil
}

// BatchCreateRadios 批量创建单选题选项
func (r *subjectInfoRepository) BatchCreateRadios(radios []*model.Radio, tx *gorm.DB) error {
	if len(radios) == 0 {
		return nil
	}
	return r.withTx(tx).Create(&radios).Error
}

// ListRadiosBySubjectID 查询单选题选项
func (r *subjectInfoRepository) ListRadiosBySubjectID(subjectID uint64, tx *gorm.DB) ([]*model.Radio, error) {
	var radios []*model.Radio
	if err := r.withTx(tx).Where("subject_id = ?", subjectID).Order("option_type").Find(&radios).Error; err != nil {
		return nil, err
	}
	return radios, nil
}

// BatchCreateMultiples 批量创建多选题选项
func (r *subjectInfoRepository) BatchCreateMultiples(multiples []*model.Multiple, tx *gorm.DB) error {
	if len(multiples) == 0 {
		return nil
	}
	return r.withTx(tx).Create(&multiples).Error
}

// ListMultiplesBySubjectID 查询多选题选项
func (r *subjectInfoRepository) ListMultiplesBySubjectID(subjectID uint64, tx *gorm.DB) ([]*model.Multiple, error) {
	var multiples []*model.Multiple
	if err := r.withTx(tx).Where("subject_id = ?", subjectID).Order("option_type").Find(&multiples).Error; err != nil {
		return nil, err
	}
	return multiples, nil
}

// CreateCategory 创建题目分类
func (r *subjectInfoRepository) CreateCategory(category *model.SubjectCategory, tx *gorm.DB) error {
	return r.withTx(tx).Create(category).Error
}

// GetCategoryByID 按 ID 查询分类
func (r *subjectInfoRepository) GetCategoryByID(id uint64, tx *gorm.DB) (*model.SubjectCategory, error) {
	var category model.SubjectCategory
	if err := r.withTx(tx).First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

// ListCategoriesByParent 按父级查询分类列表，parentID 为空则返回顶级
func (r *subjectInfoRepository) ListCategoriesByParent(parentID *uint64, tx *gorm.DB) ([]*model.SubjectCategory, error) {
	var categories []*model.SubjectCategory
	query := r.withTx(tx)
	if parentID == nil {
		query = query.Where("parent_id IS NULL")
	} else {
		query = query.Where("parent_id = ?", parentID)
	}
	if err := query.Order("id").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// CreateLabel 创建题目标签
func (r *subjectInfoRepository) CreateLabel(label *model.SubjectLabel, tx *gorm.DB) error {
	return r.withTx(tx).Create(label).Error
}

// GetLabelByID 按 ID 查询标签
func (r *subjectInfoRepository) GetLabelByID(id uint64, tx *gorm.DB) (*model.SubjectLabel, error) {
	var label model.SubjectLabel
	if err := r.withTx(tx).First(&label, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &label, nil
}

// ListLabels 查询所有标签
func (r *subjectInfoRepository) ListLabels(tx *gorm.DB) ([]*model.SubjectLabel, error) {
	var labels []*model.SubjectLabel
	if err := r.withTx(tx).Order("sort_num, id").Find(&labels).Error; err != nil {
		return nil, err
	}
	return labels, nil
}

// CreateMapping 创建题目与分类、标签关系
func (r *subjectInfoRepository) CreateMapping(mapping *model.SubjectMapping, tx *gorm.DB) error {
	return r.withTx(tx).Create(mapping).Error
}

// ListMappingsBySubjectID 按题目查询关系
func (r *subjectInfoRepository) ListMappingsBySubjectID(subjectID uint64, tx *gorm.DB) ([]*model.SubjectMapping, error) {
	var mappings []*model.SubjectMapping
	if err := r.withTx(tx).Where("subject_id = ?", subjectID).Find(&mappings).Error; err != nil {
		return nil, err
	}
	return mappings, nil
}

// DeleteMappingsBySubjectID 删除题目关联的所有关系
func (r *subjectInfoRepository) DeleteMappingsBySubjectID(subjectID uint64, tx *gorm.DB) error {
	return r.withTx(tx).Where("subject_id = ?", subjectID).Delete(&model.SubjectMapping{}).Error
}

// CreateLiked 创建点赞记录
func (r *subjectInfoRepository) CreateLiked(liked *model.SubjectLiked, tx *gorm.DB) error {
	return r.withTx(tx).Create(liked).Error
}

// GetLiked 查询点赞状态
func (r *subjectInfoRepository) GetLiked(subjectID uint64, likeUserID string, tx *gorm.DB) (*model.SubjectLiked, error) {
	var liked model.SubjectLiked
	if err := r.withTx(tx).Where("subject_id = ? AND like_user_id = ?", subjectID, likeUserID).First(&liked).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &liked, nil
}

// UpdateLikedStatus 更新点赞状态
func (r *subjectInfoRepository) UpdateLikedStatus(subjectID uint64, likeUserID string, status int, tx *gorm.DB) error {
	return r.withTx(tx).Model(&model.SubjectLiked{}).
		Where("subject_id = ? AND like_user_id = ?", subjectID, likeUserID).
		Update("is_liked", status).Error
}

// withTx 如果提供事务则使用事务，否则使用默认数据库连接
func (r *subjectInfoRepository) withTx(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return r.db
}
