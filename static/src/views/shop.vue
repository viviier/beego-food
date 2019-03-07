<template>
  <div class="root">
    <beego-header v-if="!isPopover"></beego-header>
    <el-row v-loading="loading">
      <el-table
        :data="shopFood"
        :height="tableHeight"
        style="width: 100%">
        <el-table-column
          prop="name"
          label="商品名称"
          width="180">
        </el-table-column>
        <el-table-column
          prop="description"
          label="商品简介"
          width="180">
        </el-table-column>
        <el-table-column
          align="right"
          label="个数"
        >
          <template slot-scope="scope">
            <el-input-number @change="countChange(scope.row, $event)" size="medium" v-model="scope.row.count" :min="1" label="个数"></el-input-number>
          </template>
        </el-table-column>
      </el-table>
      <div class="order">
        <el-button @click="postOrder" type="primary">确认购买</el-button>
      </div>
    </el-row>
  </div>
</template>

<script>
export default {
  props: {
    isPopover: Boolean
  },
  data () {
    return {
      shopItems: [],
      loading: true
    }
  },
  computed: {
    shopFood () {
      return this.$store.getters.foodItems
    },
    tableHeight () {
      if (this.isPopover) {
        return '350'
      }
      return ''
    }
  },
  methods: {
    countChange (item, count) {
      item.count = count
      this.$store.dispatch('changeFoodShop', item)
    },
    postOrder () {
      if (this.shopFood.length === 0) {
        return
      }
      this.$message('购买成功！')
      this.$store.dispatch('clearFoodShop')
      // 提交订单
    }
  },
  mounted () {
    this.loading = false
  }
}
</script>

<style lang="scss" scoped>
.root {
  .el-row {
    padding: 20px;

    .order {
      display: flex;
      margin-top: 40px;
      justify-content: flex-end;
    }
  }
}
</style>
