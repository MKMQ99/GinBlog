import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import Admin from '../views/Admin.vue'

// 页面路由组件
const Index = () => import(/* webpackChunkName: "Index" */ '../components/admin/Index.vue')
const AddArt = () => import(/* webpackChunkName: "AddArt" */ '../components/article/AddArt.vue')
const ArtList = () => import(/* webpackChunkName: "ArtList" */ '../components/article/ArtList.vue')
const CateList = () => import(/* webpackChunkName: "CateList" */ '../components/category/CateList.vue')
const UserList = () => import(/* webpackChunkName: "UserList" */ '../components/user/UserList.vue')

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'login',
    meta: {
      title: '请登录'
    },
    component: Login
  },
  {
    path: '/admin',
    name: 'admin',
    meta: {
      title: 'GinBlog 后台管理页面'
    },
    component: Admin,
    children: [
      {
        path: 'index',
        component: Index,
        meta: {
          title: 'GinBlog 后台管理页面'
        }
      },
      {
        path: 'addart',
        component: AddArt,
        meta: {
          title: '新增文章'
        }
      },
      {
        path: 'addart/:id',
        component: AddArt,
        meta: {
          title: '编辑文章'
        },
        props: true
      },
      {
        path: 'artlist',
        component: ArtList,
        meta: {
          title: '文章列表'
        }
      },
      {
        path: 'catelist',
        component: CateList,
        meta: {
          title: '分类列表'
        }
      },
      {
        path: 'userlist',
        component: UserList,
        meta: {
          title: '用户列表'
        }
      }
    ]
  }
]

const router = new VueRouter({
  routes
})

export default router
