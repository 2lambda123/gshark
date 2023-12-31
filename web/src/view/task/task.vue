<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增任务</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table
        :data="tableData"
        border
        ref="multipleTable"
        stripe
        style="width: 100%"
        tooltip-effect="dark"
    >
      <el-table-column
          label="类型"
          prop="taskType"
      ></el-table-column>

      <el-table-column
          label="名称"
          prop="name"
      ></el-table-column>

      <el-table-column label="创建日期">
        <template slot-scope="scope">{{
            scope.row.CreatedAt | formatDate
          }}</template>
      </el-table-column>

      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button
              class="table-button"
              @click="switchStatus(scope.row.ID, true)"
              size="small"
              type="primary"
              v-if="!scope.row.taskStatus"
          >启动</el-button
          >
          <el-button
              type="danger"
              size="mini"
              v-else
              @click="switchStatus(scope.row.ID, false)"
          >停止</el-button
          >
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[10, 30, 50, 100]"
        :style="{ float: 'right', padding: '20px' }"
        :total="total"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
        layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog
        :before-close="closeDialog"
        :visible.sync="dialogFormVisible"
        title="新增任务"
    >
      <el-form :model="formData" label-position="right" label-width="100px">
        <el-form-item label="任务类型:" required>
          <el-radio-group v-model="formData.taskType">
            <el-radio v-for="ruleType in types" :label="ruleType" :key="ruleType">{{ ruleType }}</el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="任务名称:" required>
          <el-input
              v-model="formData.name"
              clearable
              placeholder="请输入任务名称"
          ></el-input>
        </el-form-item>

        <el-form-item label="状态:">
          <el-switch v-model="formData.taskStatus"></el-switch>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="createTask" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
import { store } from '@/store/index';
import {createTask, getTaskList, switchTaskStatus} from "@/api/task";

export default {
  name: "Rule",
  mixins: [infoList],
  data() {
    return {
      listApi: getTaskList,
      dialogFormVisible: false,
      dialogBatchRules: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        name: "",
        taskType: "",
        taskStatus: true
      },
      batchRulesForm: {
        type: [],
        contents: ''
      },
      statusOptions: [
        {
          label: "disabled",
          value: 0,
        },
        {
          label: "enabled",
          value: 1,
        },
      ],
      types: ['github', 'gitlab', 'searchcode', 'domain', 'postman'],
      typeOptions: [
        {
          label: "github",
          value: "github",
        },
        {
          label: "gitlab",
          value: "gitlab",
        },
        {
          label: "searchcode",
          value: "searchcode",
        },
        {
          label: "postman",
          value: "postman"
        }
      ],
    };
  },
  filters: {
    formatDate: function (time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function (bool) {
      if (bool != null) {
        return bool ? "是" : "否";
      } else {
        return "";
      }
    },
    formatStatus: function (status) {
      const statusMap = {
        0: "disabled",
        1: "enabled",
      };
      return statusMap[status];
    },
  },
  methods: {
    async createTask() {
      const res = await createTask(this.formData);
      if (res.code === 0) {
        this.$message({
          type: "success",
          message: "创建任务成功"
        });
        this.closeDialog();
        await this.getTableData();
      }
    },
    async switchStatus(id, status) {
      status = status ? 1 : 0;
      const data = {
        id,
        status,
      };
      const res = await switchTaskStatus(data);
      if (res) {
        await this.getTableData();
      }
    },
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        name: "",
        taskType: "",
        taskStatus: 0,
      };
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
  },
};
</script>

<style>
</style>
