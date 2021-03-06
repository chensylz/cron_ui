package cron_ui

const template = `<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="UTF-8">
    <title>任务调度</title>
    <script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script>
    <link href="https://unpkg.com/element-ui/lib/theme-chalk/index.css" rel="stylesheet">
    <script src="https://unpkg.com/element-ui/lib/index.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/axios/dist/axios.min.js"></script>
    <style>
        body, div {
            margin: auto;
            padding: 0;
        }

        #web_app {
            padding: 0 10%;
        }

        a {
            text-decoration-line: none;
        }
    </style>
</head>
<body>
<div id="web_app">
    <div>
        <h1>定时任务管理</h1>
        <span style="color: gray">zaihui</span>
    </div>
    <el-divider></el-divider>
    <el-table
            :data="jobObject.jobs"
            style="width: 100%">
        <el-table-column
                label="任务ID"
                prop="id"
                min-width="180">
        </el-table-column>
        <el-table-column
                label="任务名称"
                prop="name"
                min-width="180">
        </el-table-column>
        <el-table-column
                label="表达式"
                width="260"
                prop="spec">
        </el-table-column>
        <el-table-column
                label="上次执行时间"
                width="180"
                prop="pre_time">
        </el-table-column>
        <el-table-column
                label="下次执行时间"
                width="180"
                prop="next_time">
        </el-table-column>
        <el-table-column
                fixed="right"
                label="操作"
                width="300">
            <template slot-scope="scope">
                <el-button style="margin-left: 0" @click="handleJobStart(scope.row)" size="small" type="primary">手动执行
                </el-button>
                <el-button slot="reference" size="small" @click="handleJobLog(scope.row)">手动执行日志</el-button>
            </template>
        </el-table-column>
    </el-table>

    <el-dialog
            title="手动执行日志"
            :visible.sync="jobObject.dialog.isShow"
            width="30%"
            :before-close="handleJobDialogClose"
            center>
        <el-table :data="jobObject.dialog.histories">
            <el-table-column property="index" label="索引" min-width="150"></el-table-column>
            <el-table-column property="time" label="执行时间" min-width="200"></el-table-column>
        </el-table>
        <span slot="footer" class="dialog-footer">
          <el-button type="primary" @click="jobObject.dialog.isShow = false">确 定</el-button>
        </span>
    </el-dialog>
</div>
</body>
<script>
    const app = new Vue({
        el: "#web_app",
        data: {
            jobObject: {
                jobs: [],
                dialog: {
                    isShow: false,
                    histories: []
                }
            },
            headerObject: {
                selectIndex: 1,
            },
        },
        created() {
            this.pageInit()
        },
        methods: {
            pageInit() {
                axios.get("/cronjob").then((success) => {
                    if (success.status === 200) {
                        this.jobObject.jobs = success.data
                    }
                })
            },
            /**
             * 处理任务对话框的回调
             */
            handleJobDialogClose() {
                this.jobObject.dialog.histories = []
                this.jobObject.dialog.isShow = false
            },
            /**
             * 处理任务编辑
             * @param job 任务
             */
            handleJobStart(job) {
                axios.post("/cronjob/" + job.id).then((success) => {
                    if (success.status === 200) {
                        this.$notify({
                            title: '成功',
                            message: '执行任务成功',
                            type: 'success'
                        });
                        this.pageInit()
                        return
                    }
                }).catch(resp => {
                    this.$notify({
                        title: '警告',
                        message: resp.response.data.message,
                        type: 'warning'
                    });
                })
            },
            handleJobLog(job) {
                this.jobObject.dialog.histories = job.manual_history ? this.toShowHistory(job.manual_history) : []
                this.jobObject.dialog.isShow = true
            },
            toShowHistory(history) {
                let result = [];
                for (let i = history.length - 1; i >= 0; i--) {
                    result.push({"index": history.length - i, "time": history[i]})
                }
                return result
            }
        }
    })
</script>
</html>
`
