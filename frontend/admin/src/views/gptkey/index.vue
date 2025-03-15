<template>
    <div class="app-container">
        <!--工具栏-->
        <div class="head-container">
            <!-- 新增 -->
        </div>
        <!--表单组件-->
        <!--表格渲染-->
        <el-table v-loading="loading" :data="data" size="small" style="width: 100%;">
            <el-table-column prop="id" label="id" width="100" />
            </el-table-column>
            <el-table-column prop="apiKey" label="API KEY" />
            </el-table-column>
            <el-table-column prop="status" label="状态" />
            </el-table-column>
            <el-table-column prop="remarks" label="说明" />
            </el-table-column>
        </el-table>
    </div>
</template>
       
<script>
import initData from '@/mixins/crud'
export default {
    components: {},
    mixins: [initData],
    created() {
        this.$nextTick(() => {
            this.init()
        })
    },
    methods: {
        beforeInit() {
            this.url = 'admin/gptkey'
            const sort = 'id,desc'
            this.params = { page: this.page, size: this.size, sort: sort, userType: this.userType }
            const query = this.query
            const type = query.type
            const value = query.value
            if (type && value) { this.params[type] = value }
            return true
        },
    },
}
</script>

<style scoped></style>
