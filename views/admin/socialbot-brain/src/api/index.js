
const api = {
    Host: 'http://localhost:8080',
    UploadSingle: '/v1/adminApi/upload/single',
    Login: '/v1/adminApi/login',
    Logout: '/v1/adminApi/logout',

    AddServerConfig: '/v1/adminApi/config/addServer',
    UpdateServerConfig: '/v1/adminApi/config/updateServer',
    ListServerConfig: '/v1/adminApi/config/listServer',
    DeleteServerConfig: '/v1/adminApi/config/deleteServer',

    AddCategory: '/v1/adminApi/category/add',
    UpdateCategory: '/v1/adminApi/category/update',
    ListCategory: '/v1/adminApi/category/list',
    DeleteCategory: '/v1/adminApi/category/delete',
    ListCategoryWithTags: '/v1/adminApi/category/listWithTags',

    AddTag: '/v1/adminApi/tag/add',
    UpdateTag: '/v1/adminApi/tag/update',
    ListTag: '/v1/adminApi/tag/list',
    DeleteTag: '/v1/adminApi/tag/delete',


    AddRobotServer: '/v1/adminApi/robot/addServer',
    DeleteRobotServer: '/v1/adminApi/robot/deleteServer',
    ListRobotServer: '/v1/adminApi/robot/listServer',

    AddCrawler: '/v1/adminApi/crawler/add',
    UpdateCrawler: '/v1/adminApi/crawler/update',
    ListCrawler: '/v1/adminApi/crawler/list',
    ListCrawlerItem: '/v1/adminApi/crawler/listItem',

    AddCommissionProduct: '/v1/adminApi/media/addCommissionProduct',
    AddSocialMediaFromCrawler: '/v1/adminApi/media/addSocialMediaFromCrawler',

};
export default api
