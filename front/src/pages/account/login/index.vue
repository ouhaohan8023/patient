<template>
    <div class="page-account login">
        <div v-if="showI18n" class="page-account-header">
            <i-header-i18n />
        </div>
        <div class="page-account-container">
            <div class="page-account-top">
                <div class="page-account-top-logo">
                    <img src="@/assets/images/logo.png" alt="logo">
                </div>
                <div class="page-account-top-desc">Done. Patient Register</div>
            </div>
          <Form ref="formValidate" :model="formItem" :rules="ruleValidate" :label-width="0">
            <FormItem prop="username">
              <Input prefix="ios-contact" v-model="formItem.username" placeholder="Username"></Input>
            </FormItem>
            <FormItem prop="username">
              <Input type="password" prefix="ios-contact" v-model="formItem.password" placeholder="Password"></Input>
            </FormItem>
            <div class="page-account-auto-login">
              <Checkbox v-model="autoLogin" size="large">{{ $t('page.login.remember') }}</Checkbox>
              <!--                    <a href="">{{ $t('page.login.forgot') }}</a>-->
            </div>
            <FormItem>
              <Button type="primary" class="ivu-btn-custom registerBtn" @click="handleSubmit('formValidate')">Login</Button>
            </FormItem>
          </Form>
        </div>
        <i-copyright />
    </div>
</template>
<style lang="less">
.login {
  .registerBtn {
    width: 100%;
    height: 40px;
    font-size: 16px;
  }
  .ivu-input {
    height: 40px;
    line-height: 40px;
  }
  .ivu-input-prefix i, .ivu-input-suffix i {
    font-size: 18px;
    line-height: 40px;
    color: #808695;
  }
}
</style>
<script>
    import iCopyright from '@/components/copyright';
    import { mapActions } from 'vuex';
    import mixins from '../mixins';

    export default {
        mixins: [ mixins ],
        components: { iCopyright },
        data () {
            return {
                autoLogin: true,
                formItem: {
                    username: 'admin',
                    password: 'admin'
                },
                ruleValidate: {
                    username: [
                        { required: true, message: 'Username is required', trigger: 'blur' }
                    ],
                    password: [
                        { required: true, message: 'Password is required', trigger: 'blur' }
                    ]
                }
            }
        },
        methods: {
            ...mapActions('admin/account', [
                'login'
            ]),
            /**
             * @description 登录
             * 表单校验已有 iView Pro 自动完成，如有需要修改，请阅读 iView Pro 文档
             */
            handleSubmit (name) {
                let _that = this
                this.$refs[name].validate((valid) => {
                    if (valid) {
                        this.login(this.formItem)
                            .then((res) => {
                                console.log(res, 99)
                                if (res.status === 200) {
                                    // 重定向对象不存在则返回顶层路径
                                    this.$router.replace(this.$route.query.redirect || '/');
                                } else {
                                    _that.$Message.error(res.msg);
                                }
                        });
                    } else {
                        this.$Message.error('Fail!');
                    }
                })
            }
        }
    };
</script>
