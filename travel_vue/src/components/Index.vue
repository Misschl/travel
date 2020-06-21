<template>
    <div class="container">
        <div class="view">
            <router-view/>
        </div>
        <el-footer>
            <el-row :gutter="20" type="flex" class="row-bg" justify="space-around">
                <el-col :span="6">
                    <div class="menu" @click="toMenuPage('/home')">
                        <i class="el-icon-s-home"/>
                        <div>首页</div>
                    </div>
                </el-col>
                <el-col :span="6">
                    <div class="menu" @click="toMenuPage('/behavior')">
                        <i class="el-icon-coffee-cup"/>
                        <div>动态</div>
                    </div>
                </el-col>
                <el-col :span="6">
                    <div class="menu" @click="toMenuPage('/friends')">
                        <el-badge :value="3" type="primary">
                            <i class="el-icon-star-off"/>
                        </el-badge>
                        <div>好友</div>
                    </div>
                </el-col>
                <el-col :span="6">
                    <div class="menu" @click="toMenuPage('/user')">
                        <i class="el-icon-s-custom"/>
                        <div>我的</div>
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
            return {}
        },
        methods: {
            toMenuPage(path) {
                this.$router.push(path)
            },
            async getCurrentUser() {
                const {data: resp} = await this.$http.get("account/info");
                console.log(resp)
            }
        }, created() {
            this.getCurrentUser()
        }
    }
</script>

<style scoped>

    .container {
        height: 100%;
    }

    .view {
        margin-top: 10px;
        margin-left: 8px;
        margin-right: 8px;
        height: 89%;
    }

    .el-footer {
        position: relative;
        bottom: 0;
        width: 100%;
        border-top: 1px solid #C0C4CC;
    }


    .menu {
        position: relative;
        color: #C0C4CC;
        text-align: center;
        padding-top: 20%;
    }


</style>
