<template>
  <div class="container">
    <div class="row">
      <div class="col-md-12">
        <div class="row">
          <div class="col card-item">
            <p>Available</p>
            <div class="border-text">0 TMUN</div>
            <button>Stake</button>
          </div>
          <div class="col card-item">
            <p>Staked</p>
            <div class="border-text">0 TMUN</div>
            <div class="row">
              <div class="col-md-6">
                <button class="border">Undelegate</button>
              </div>
              <div class="col-md-6">
                <button class="border col-md-6">Redelegate</button>
              </div>
            </div>
          </div>
          <div class="col card-item">
            <p>Rewards</p>
            <div class="border-text">0 TMUN</div>
            <button>Claim</button>
          </div>
          <div class="col card-item">
            <p>Unbonding</p>
            <div class="border-text">0 TMUN</div>
            <div class="transparent">Unbonding period: 14 days</div>
          </div>
          <div class="col card-item">
            <p>Staking APR</p>
            <div class="border-text">86.87%</div>
          </div>
        </div>
      </div>
      <div class="col-md-12 validators-box">
        <div class="table-badge">
          <div>All Validators (154)</div>
          <div>
            <button>Active</button>
            <button class="border">Inactive</button>
          </div>
        </div>
        <table>
          <thead>
            <tr>
              <td width="50%"><a>NAME</a></td>
              <td class="text-center"><a>VOTING POWER</a></td>
              <td class="text-center"><a>COMMISSION</a></td>
              <td class="text-center">ACTIONS</td>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>
                <span class="avatar-name">web34ever</span>
              </td>
              <td class="text-center">1,908,601.47 TMUN</td>
              <td class="text-center">5%</td>
              <td class="text-center">
                <a>Manage →</a>
              </td>
            </tr>
            <tr>
              <td>
                <span class="avatar-name">web34ever</span>
              </td>
              <td class="text-center">1,908,601.47 TMUN</td>
              <td class="text-center">5%</td>
              <td class="text-center">
                <a>Manage →</a>
              </td>
            </tr>
            <tr>
              <td>
                <span class="avatar-name">web34ever</span>
              </td>
              <td class="text-center">1,908,601.47 TMUN</td>
              <td class="text-center">5%</td>
              <td class="text-center">
                <a>Manage →</a>
              </td>
            </tr>
            <tr>
              <td>
                <span class="avatar-name">web34ever</span>
              </td>
              <td class="text-center">1,908,601.47 TMUN</td>
              <td class="text-center">5%</td>
              <td class="text-center">
                <a>Manage →</a>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'

export default {
  name: 'Portfolio',

  setup() {
    // store
    let $s = useStore()

    // computed
    let address = computed(() => $s.getters['common/wallet/address'])

    return {
      address
    }
  }
}
</script>

<style scoped>
.row {
  display: flex;
  flex-wrap: wrap;
}

.col {
  flex-grow: 1;
  padding: 20px;
}

.card-item {
  padding: 8px;
}

.card-item p {
  font-size: 16px;
  font-weight: 500;
  text-align: center;
}

.card-item .border-text {
  border: 1px solid #ddd;
  border-radius: 1rem;
  font-size: 20px;
  font-weight: 800;
  padding: 12px 0;
  text-align: center;
}

.card-item button {
  background: #db2777;
  border: none;
  border-radius: 1rem;
  color: white;
  cursor: pointer;
  font-size: 16px;
  font-weight: 500;
  margin-top: 10px;
  padding: 12px 0;
  width: 100%;
  text-align: center;
}

.card-item button:hover {
  opacity: 0.8;
}

.card-item button.border {
  border: 1px solid #db2777;
  background: transparent;
  color: #db2777;
  padding: 10px 0;
}

.card-item button.border:hover {
  background: #db2777;
  color: white;
  opacity: 1;
}

.card-item .transparent {
  width: 100%;
  font-size: 14px;
  font-weight: 500;
  margin-top: 10px;
  padding: 12px 0;
  text-align: center;
}

.validators-box {}

.validators-box table {
  width: 100%;
}

.text-center {
  text-align: center;
}

.table-badge {
  align-items: center;
  border: 1px solid #ddd;
  border-top-left-radius: 1rem;
  border-top-right-radius: 1rem;
  display: flex;
  font-size: 20px;
  justify-content: space-between;
  margin-top: 30px;
  padding: 16px;
}

.table-badge button {
  font-size: 14px;
  background: #db2777;
  border-radius: 1rem;
  border: 1px solid #db2777;
  color: white;
  cursor: pointer;
  font-weight: 800;
  padding: 10px 20px;
  margin-left: 16px;
}

.table-badge button:hover {
  opacity: 0.8;
}

.table-badge button.border {
  background: transparent;
  color: #db2777;
}

.table-badge button.border:hover {
  background: #db2777;
  color: white;
  font-weight: 800;
  opacity: 1;
}

table
{
  font-size: 14px;
  font-weight: 500;
  border-spacing: 0px;
  border-left: 1px solid #ddd;
  border-right: 1px solid #ddd;
}

table td {
  padding: 10px;
  border: none;
  border-bottom: 1px solid #ddd;
}

table tbody td {
  padding: 16px 10px;
}

table a {
  color: #db2777;
  font-weight: 700;
  cursor: pointer;
}

table a:hover {
  opacity: 0.8;
}
</style>
