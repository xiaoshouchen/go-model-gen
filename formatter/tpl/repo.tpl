type %sRepo struct{
    db *gorm.DB
}

func New%sRepo(db *gorm.DB)*%sRepo{
    return &%sRepo{db: db}
}
