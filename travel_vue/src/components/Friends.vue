<template>
    <div class="friend-container">
        <div class="head"></div>
        <div class="search_bar">
            <el-autocomplete v-model="query" :fetch-suggestions="querySearchAsync" placeholder="请输入内容"
                             @select="handleSelect">
                <i class="el-icon-edit el-input__icon" slot="suffix"/>
            </el-autocomplete>
        </div>
        <div class="menu-tab">
            <el-tabs type="card" stretch>

                <el-tab-pane label="好友列表" name="first">
                    <!--好友布局-->
                    <el-col class="friend-item">
                        <el-col>
                            <el-image src="https://ivy-static.oss-cn-hangzhou.aliyuncs.com/avatar.jpg"  style="width: 40px; height:40px"></el-image>
<!--                            <el-avatar src="https://ivy-static.oss-cn-hangzhou.aliyuncs.com/avatar.jpg"/>-->
                        </el-col>
                        <el-col class="username">
                            万总
                        </el-col>
                        <el-col class="sign">
                            万总签名
                        </el-col>
                    </el-col>
                    <el-col class="friend-item">
                        <el-col>
                            <el-image src="https://ivy-static.oss-cn-hangzhou.aliyuncs.com/avatar.jpg"  style="width: 40px; height:40px"></el-image>
                            <!--                            <el-avatar src="https://ivy-static.oss-cn-hangzhou.aliyuncs.com/avatar.jpg"/>-->
                        </el-col>
                        <el-col class="username">
                            万总
                        </el-col>
                        <el-col class="sign">
                            万总签名
                        </el-col>
                    </el-col>
                    <el-col class="friend-item">
                        <el-col>
                            <el-image src="https://ivy-static.oss-cn-hangzhou.aliyuncs.com/avatar.jpg"  style="width: 40px; height:40px"></el-image>
                            <!--                            <el-avatar src="https://ivy-static.oss-cn-hangzhou.aliyuncs.com/avatar.jpg"/>-->
                        </el-col>
                        <el-col class="username">
                            万总
                        </el-col>
                        <el-col class="sign">
                            万总签名
                        </el-col>
                    </el-col>
                    <el-col class="friend-item">
                        <el-col>
                            <el-image src="https://ivy-static.oss-cn-hangzhou.aliyuncs.com/avatar.jpg"  style="width: 40px; height:40px"></el-image>
                            <!--                            <el-avatar src="https://ivy-static.oss-cn-hangzhou.aliyuncs.com/avatar.jpg"/>-->
                        </el-col>
                        <el-col class="username">
                            万总
                        </el-col>
                        <el-col class="sign">
                            万总签名
                        </el-col>
                    </el-col>
                </el-tab-pane>


                <el-tab-pane label="消息" name="second">配置管理</el-tab-pane>

            </el-tabs>
        </div>
    </div>
</template>

<script>
    export default {
        name: "Friends",
        data() {
            return {
                query: "",
                friendList: [],
                num: [1, 2, 3, 4]
            }
        },
        methods: {
            querySearchAsync() {
                console.log(this.query)
            },
            handleSelect() {

            },
            count() {

            }

        }
    }
</script>

<style scoped>
    .friend-container {
        height: 100%;
    }

    .head {
        height: 5%;
    }

    .el-autocomplete {
        width: 100%;
    }

    .menu-tab {
        padding-top: 10px;
    }

    .friend-item {
        margin-top: 10px;
        height: 50px;
        /*border-top: 1px solid #C0C4CC;*/
        border-radius: 4px
    }


    .el-col {
        position: relative;
    }

    .username {
        position: absolute;
        left: 50px;
        font-size: 20px;
    }

    .sign {
        position: absolute;
        left: 50px;
        top: 20px;
    }

</style>
