<template>
    <div>
        <router-view/>
        <el-footer>
            <el-row :gutter="20" type="flex" class="row-bg" justify="center">
                <el-col :span="6" v-for="menu in menuList" :key="menu.path">
                    <div class="menu" @click="toMenuPage(menu.path)">
                        <i :class="menu.icon"/>
                        <div>{{menu.name}}</div>
                    </div>
                </el-col>
            </el-row>
        </el-footer>
    </div>
</template>

<script>
    export default {
        name: "Index",
        data() {
            return {
                menuList: [
                    {name: "首页", icon: "el-icon-s-home", path: "/home"},
                    {name: "首页", icon: "el-icon-s-home", path: "/home"},
                    {name: "首页", icon: "el-icon-s-home", path: "/home"},
                    {name: "我的", icon: "el-icon-s-custom", path: "/user"},
                ]
            }
        },
        methods: {
            toMenuPage(path) {
                this.$router.push(path)
            }
        }
    }
</script>

<style scoped>
    .el-footer {
        position: absolute;
        bottom: 0;
        width: 100%;
        border-top: 1px solid #C0C4CC;
        min-width: 100%;
    }

    .menu {
        color: #C0C4CC;
        text-align: center;
        margin-top: 8px;
        margin-right: 20px;
    }
</style>
