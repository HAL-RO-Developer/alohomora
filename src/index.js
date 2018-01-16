import Vue from 'vue'
import VueRouter from 'vue-router'
import Index from './components/index/index.vue'
import Product from './components/product/product.vue'
import NotFound from './components/notFound/notFound.vue'

Vue.component('index', Index)

Vue.use(VueRouter)

const routes = [
    { path: "/", component: Index },
    { path: "/m*", component: Product },
    { path: "*", component: NotFound },
]
const router = new VueRouter({  mode: 'history', routes })

new Vue({
    router
}).$mount("#app")