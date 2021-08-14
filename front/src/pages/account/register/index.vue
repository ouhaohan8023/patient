<template>
    <div class="page-account register">
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
              <Input prefix="ios-contact" v-model="formItem.username" placeholder="Name"></Input>
            </FormItem>
            <FormItem prop="birthday">
              <DatePicker type="date" placeholder="Date Of Birth" :start-date="new Date(2000, 0, 1)" style="width: 100%" v-model="formItem.birthday"></DatePicker>
            </FormItem>
            <FormItem prop="mobile">
              <Input prefix="md-phone-portrait" v-model="formItem.mobile" placeholder="Phone Number"></Input>
            </FormItem>
            <FormItem prop="email">
              <Input prefix="md-mail" v-model="formItem.email" placeholder="Email"></Input>
            </FormItem>
            <FormItem prop="address">
              <Input prefix="md-locate" v-model="formItem.address" placeholder="Address"></Input>
            </FormItem>
            <FormItem prop="appointment">
              <DatePicker type="datetime" placeholder="Appointment Time" style="width: 100%" v-model="formItem.appointment"></DatePicker>
            </FormItem>
            <FormItem>
              <div class="demo-upload-list" v-for="(item, index) in uploadList" v-bind:key="index">
                <template v-if="item.status === 'finished'">
                  <img :src="item.url">
                  <div class="demo-upload-list-cover">
                    <Icon type="ios-trash-outline" @click.native="handleRemove(item)"></Icon>
                  </div>
                </template>
                <template v-else>
                  <Progress v-if="item.showProgress" :percent="item.percentage" hide-info></Progress>
                </template>
              </div>
              <Upload
                    ref="upload"
                    :show-upload-list="false"
                    :default-file-list="defaultList"
                    :on-success="handleSuccess"
                    :format="['jpg','jpeg','png']"
                    :max-size="2048"
                    :on-format-error="handleFormatError"
                    :on-exceeded-size="handleMaxSize"
                    :before-upload="handleBeforeUpload"
                    multiple
                    type="drag"
                    action="/api/upload"
                    style="width:100%;"
                    v-bind:style="uploadBtnStyleObject"
                >
                  <div style="width: 100%;height:58px;line-height: 58px;">
                    <Icon type="ios-camera" size="20"></Icon>
                    Please Drag Or Select Your Drive Licence Here
                  </div>
                </Upload>
            </FormItem>

            <FormItem>
              <Button class="ivu-btn-custom registerBtn" @click="handleSubmit('formValidate')">Register</Button>
            </FormItem>
          </Form>
        </div>
        <i-copyright />
    </div>
</template>
<style lang="less">
.register .ivu-input-prefix i, .ivu-input-suffix i {
  font-size: 18px;
  line-height: 40px;
  color: #808695;
}
.register .ivu-input {
  height: 40px;
}
.register .ivu-input:hover {
  border-color: #fd2d55;
}
.register .registerBtn {
  width: 100%;
  height: 40px;
  font-size: 16px;
}
.register .ivu-btn-custom {
  background-color: #fd2d55;
  color: #fffdef;
  border: 1px solid transparent;
  border-radius: 0.3125rem;
  transition: color 200ms ease, background-color 200ms ease, opacity 200ms ease, border 200ms ease;
  line-height: 1.5;
  font-weight: 500;
  text-align: center;
  cursor: pointer;
}
.register .demo-upload-list{
  display: inline-block;
  width: 100%;
  height: 250px;
  text-align: center;
  line-height: 60px;
  border: 1px solid transparent;
  border-radius: 4px;
  overflow: hidden;
  background: #fff;
  position: relative;
  box-shadow: 0 1px 1px rgba(0,0,0,.2);
  margin-right: 4px;
}
.register .demo-upload-list img{
  width: 100%;
  height: 100%;
}
.register .demo-upload-list-cover{
  display: none;
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  background: rgba(0,0,0,.6);
}
.register .demo-upload-list:hover .demo-upload-list-cover{
  display: block;
}
.register .demo-upload-list-cover i{
  color: #fff;
  font-size: 20px;
  cursor: pointer;
  margin: 0 2px;
}
.register .ivu-btn-custom:hover {
  background-color: #FD0736;
  border: 1px solid transparent;
  color: #fffdef;
}

</style>
<script>
    import iCopyright from '@/components/copyright';
    import { mapActions } from 'vuex';
    import mixins from '../mixins';
    import { AccountRegister } from '../../../api/account';

    export default {
        mixins: [ mixins ],
        components: { iCopyright },
        data () {
            return {
                birthday: '',
                address: '',
                defaultList: [
                ],
                imgName: '',
                visible: false,
                uploadList: [],
                displayUpload: true,
                uploadBtnStyleObject: {
                    display: 'inline-block'
                },
                formItem: {
                    username: '',
                    birthday: '',
                    mobile: '',
                    email: '',
                    address: '',
                    appointment: '',
                    img: ''
                },
                ruleValidate: {
                    username: [
                        { required: true, message: 'The Name Cannot Be Empty', trigger: 'blur' }
                    ],
                    mobile: [
                        { required: true, message: 'Phone Cannot Be Empty', trigger: 'blur' }
                    ],
                    email: [
                        { required: true, message: 'Email Address Cannot Be Empty', trigger: 'blur' },
                        { type: 'email', message: 'Incorrect Email Format', trigger: 'blur' }
                    ],
                    address: [
                        { required: true, message: 'The Address Cannot Be Empty', trigger: 'blur' }
                    ],
                    birthday: [
                        { required: true, type: 'date', message: 'Please Select The Birthday', trigger: 'change' }
                    ],
                    appointment: [
                        { required: true, type: 'date', message: 'Please Select The Appointment Time', trigger: 'change' }
                    ]
                }
            }
        },
        computed: {
        },
        methods: {
            ...mapActions('admin/account', [
                'register',
                'login'
            ]),
            handleChangePassword (val) {
                this.passwordLen = val.length;
            },
            /**
             * @description 注册
             * 表单校验已有 iView Pro 自动完成，如有需要修改，请阅读 iView Pro 文档
             */
            handleSubmit (name) {
                this.$refs[name].validate((valid) => {
                    if (valid) {
                        let _that = this
                        AccountRegister(this.formItem).then(function (res) {
                            // console.log(res)
                            if (res.status === 500) {
                                if (res.content) {
                                    _that.$Message.error(res.content.msg);
                                } else {
                                    _that.$Message.error(res.msg);
                                }
                            } else {
                                _that.$Message.success('Register Success!');
                                _that.$router.replace({ name: 'register-result' });
                            }
                        })
                        // this.$Message.success('Success!');
                    } else {
                        // this.$Message.error('Fail!');
                    }
                })
            },
            /**
             * @description get sms code
             * */
            handleGetCaptcha () {

            },
            handleRemove (file) {
                const fileList = this.$refs.upload.fileList;
                this.$refs.upload.fileList.splice(fileList.indexOf(file), 1);
                this.uploadBtnStyleObject.display = 'inline-block'
            },
            handleSuccess (res, file) {
                file.url = res.content.url;
                this.formItem.photo = file.url
                this.uploadBtnStyleObject.display = 'none'
                console.log(this.uploadBtn)
            },
            handleFormatError (file) {
                this.$Notice.warning({
                    title: 'The file format is incorrect',
                    desc: 'File format of ' + file.name + ' is incorrect, please select jpg or png.'
                });
            },
            handleMaxSize (file) {
                this.$Notice.warning({
                    title: 'Exceeding file size limit',
                    desc: 'File  ' + file.name + ' is too large, no more than 2M.'
                });
            },
            handleBeforeUpload () {
                const check = this.uploadList.length < 1;
                if (!check) {
                    this.$Notice.warning({
                        title: 'Up to One pictures can be uploaded.'
                    });
                }
                return check;
            }
        },
        mounted () {
            this.uploadList = this.$refs.upload.fileList;
        }
    };
</script>
