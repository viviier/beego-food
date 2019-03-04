<template>
  <div class="root" @scroll="getScroll">
    <h2>Beego-food</h2>
    <el-row v-loading="loading">
      <el-col :span="4" v-for="item in foodlist" :key="item.index" :offset="0">
        <el-card shadow="hover" :body-style="{ padding: '0px' }">
          <img :src="item.url" class="image" @click="foodDetail(item)">
          <div style="padding: 14px;">
            <span class="name">{{ item.name }}</span>
            <div class="bottom">
              <span class="description" :title="item.description">{{ item.description || 'no description' }}</span>
              <div>
                <i class="el-icon-info"></i>
                <i class="el-icon-star-off"></i>
                <i class="el-icon-goods"></i>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
    <el-dialog :visible.sync="dialogTableVisible" width="80%">
      <div class="dialog-content">
        <div class="info">
          <img :src="detailData.url" />
          <h4>{{ detailData.name }}</h4>
          Tags: &nbsp;&nbsp;<el-tag size="small" type="info">{{ detailData.role ? '咖啡' : '蛋糕'}}</el-tag>
        </div>
        <div class="more">
          <h4>description: </h4>
          <span>{{ detailData.description || 'no description' }}</span>
          <div>
            <el-button icon="el-icon-star-off" circle></el-button>
            <el-button type="primary" size="medium" @click="addFoodToShop(detailData)">加入购物车<i class="el-icon-goods el-icon--right"></i></el-button>
          </div>
        </div>
      </div>
    </el-dialog>
    <el-popover
      placement="top"
      width="800"
      trigger="click"
    >
      <el-table :data="shopFood" style="width: 100%" height="700">
        <el-table-column width="180%" property="name" label="名称"></el-table-column>
        <el-table-column width="180" property="id" label="个数"></el-table-column>
        <el-table-column property="description" label="简介"></el-table-column>
      </el-table>
      <el-button slot="reference" type="primary" icon="el-icon-goods" circle class="shop-button"></el-button>
    </el-popover>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data () {
    return {
      foodlist: [],
      detailData: {},
      dialogTableVisible: false,
      loading: true,
      currentDate: new Date()
    }
  },
  methods: {
    getFoodList () {
      axios
        .get('/api/fd.items.list')
        .then(res => {
          this.foodlist = this.foodlist.concat(res.data)
        })
    },
    foodDetail (item) {
      console.log('?')
      this.detailData = item
      this.dialogTableVisible = true
    },
    getScroll () {
      console.log('ok')
    },
    addFoodToShop (item) {
      this.$store.dispatch('addFoodShop', item)
    }
  },
  computed: {
    shopFood () {
      return this.$store.getters.foodItems
    }
  },
  mounted () {
    this.getFoodList()
    window.addEventListener('scroll', () => {
      let pageHeight = Math.max(document.body.scrollHeight, document.body.offsetHeight)
      let viewportHeight = window.innerHeight ||
          document.documentElement.clientHeight ||
          document.body.clientHeight || 0
      let scrollHeight = window.pageYOffset ||
          document.documentElement.scrollTop ||
          document.body.scrollTop || 0
      if (pageHeight - viewportHeight - scrollHeight < 20) {
        this.getFoodList()
      }
    })
    this.loading = false
  }
}
</script>

<style lang="scss" scoped>
.root {
  h2 {
    font-size: 36px;
    margin-left: 40px;
  }

  .el-row {
    padding: 20px;

    .el-card {
      margin: 20px 20px 10px;

      img {
        width: 100%;
        height: 100%;
        display: block;
      }

      span {
        display: block;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
      }

      div {
        .name {
          font-size: 12px;
        }

        .description {
          margin-top: 5px;
        }

        .bottom {
          div {
            display: flex;
            margin-top: 10px;

            i {
              margin-right: 5px;

              &:first-child {
                flex: 1;
                margin-right: 0;
              }
            }
          }
        }
      }
    }
  }

  .dialog-content {
    display: flex;

    .info {
      flex: 1;
    }

    .more {
      flex: 2;
      display: flex;
      flex-direction: column;

      div {
        flex: 1;
        display: flex;
        justify-content: flex-end;
        align-items: flex-end;
      }
    }
  }

  .shop-button {
    position: fixed;
    right: 50px;
    bottom: 50px;
  }
}

</style>
