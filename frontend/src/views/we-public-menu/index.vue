<template>
  <div style="display: block; flex: 1; height: 100%">
    <a-row>
      <a-col :span="24">
        <div class="header-box">
          <div class="info_tips_wrap">
            <div class="icon suc_icon" />
          </div>
          <div class="content">
            <div class="in-content">
              <p class="title">菜单已发布</p>
              <p class="title"
                >可在手机查看菜单内容，若尚未生效，请稍后查看。若停用菜单，
                <a
                  style="color: rgba(13, 132, 255, 0.79)"
                  href="https://mp.weixin.qq.com/"
                  target="_blank"
                  >请前往微信官方公众平台</a
                ></p
              >
            </div>
          </div>
        </div>
      </a-col>
    </a-row>
    <a-row>
      <a-col :span="24">
        <div class="container-box">
          <!--      左侧菜单编辑器-->
          <div class="phone-box">
            <div class="menu__preview-hd" />
            <div class="menu__preview-bd">
              <div class="menu__preview-bottom">
                <div class="menu__keyboard">
                  <div class="menu__keyboard-icon" />
                </div>
                <div class="menu-list">
                  <a-button
                    v-if="menuData.button.length <= 0"
                    type="text"
                    class="menu-add-no-menu"
                    @click="addMainMenu"
                  >
                    <div class="add-no-menu">
                      <div class="menu__add-icon" />
                      <span>添加菜单</span>
                    </div>
                  </a-button>
                  <div
                    v-for="item in menuData.button"
                    v-else
                    :key="item.id"
                    :class="`menu-list-item ${
                      selectMenuId === item.id && selectMenuSubId === 0
                        ? 'menu-box-shadow'
                        : ''
                    }`"
                  >
                    <div class="menu__preview-line" />
                    <div
                      :class="`menu-item ${
                        selectMenuId === item.id && selectMenuSubId === 0
                          ? 'menu-item-color'
                          : ''
                      }`"
                      @click="selectMenu(item.id)"
                      >{{ item.name }}
                    </div>
                    <!--                  二级菜单-->
                    <div v-if="selectMenuId === item.id" class="submenu">
                      <!--                    二级-->
                      <div
                        v-for="sub_item in item.subButton"
                        :key="sub_item.id"
                        :class="`menu-item-sub ${
                          selectMenuSubId === sub_item.id
                            ? 'menu-box-shadow menu-item-color'
                            : ''
                        } ${
                          selectMenuSubId !== sub_item.id
                            ? 'menu-item-border-color'
                            : ''
                        }`"
                      >
                        <span
                          style="
                            flex: 1;
                            height: 100%;
                            display: flex;
                            justify-content: center;
                            align-items: center;
                          "
                          @click="selectSubMenu(sub_item.id)"
                          >{{ sub_item.name }}</span
                        >
                        <div class="sub-menu-bar">
                          <div
                            v-if="selectMenuSubId === sub_item.id"
                            class="sub-menu-bar-box"
                          >
                            <div
                              class="sub-center-bar bar-hover"
                              @click="onDeleteSub"
                            />
                          </div>
                        </div>
                      </div>
                      <!--                    二级增加按钮-->
                      <div
                        v-if="item.subButton.length < 5"
                        class="add-button-sub"
                      >
                        <a-button
                          type="text"
                          title="最多添加5个子菜单"
                          @click="addSubMenu"
                        >
                          <template #icon>
                            <icon-plus />
                          </template>
                          <!-- Use the default slot to avoid extra spaces -->
                          <template #default>添加</template>
                        </a-button>
                      </div>
                      <i
                        :class="`arrow arrow_out ${
                          selectMenuSubId === 5 && item.subButton.length === 5
                            ? 'arrow_out_select'
                            : ''
                        }`"
                      />
                      <i
                        :class="`arrow arrow_in ${
                          selectMenuSubId === 5 && item.subButton.length === 5
                            ? 'arrow_in_select'
                            : ''
                        }`"
                      />
                    </div>
                    <div class="menu-bar">
                      <div
                        v-if="selectMenuId === item.id"
                        :class="`menu-bar-box ${
                          menuData.button.length > 1 ? 'bar-padding' : ''
                        }`"
                      >
                        <div
                          v-if="item.id !== menuData.button[0].id"
                          class="left-bar bar-hover"
                          @click="onLeftMove"
                        />
                        <div class="center-bar bar-hover" @click="onDelete" />
                        <div
                          v-if="
                            menuData.button[1] &&
                            item.id !== menuData.button[1].id &&
                            menuData.button.length === 2
                          "
                          class="right-bar bar-hover"
                          @click="onRightMove"
                        />
                        <div
                          v-if="
                            menuData.button.length > 2 &&
                            item.id !== menuData.button[2].id
                          "
                          class="right-bar bar-hover"
                          @click="onRightMove"
                        />
                      </div>
                    </div>
                  </div>
                  <!--                  主菜单添加按钮-->
                  <div
                    v-if="
                      menuData.button.length < 3 && menuData.button.length > 0
                    "
                    class="add-button"
                  >
                    <div class="menu_button_preview-line" />
                    <a-button type="text" @click="addMainMenu">
                      <i class="icon14_menu_add" />
                    </a-button>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <!--        右侧编辑器-->
          <div class="right-editor-box">
            <div v-if="menuData.button.length > 0" class="attribute-box">
              <!--              没有有二级菜单-->
              <div
                v-if="
                  (!isHaveSubMenu() && selectMenuId !== 0) ||
                  selectMenuSubId !== 0
                "
                class="custom-no-have-menu-box"
              >
                <h3 style="padding-bottom: 25px">菜单信息</h3>
                <a-form
                  ref="formVisible"
                  :model="formData"
                  :rules="rule"
                  :style="{ width: '600px' }"
                >
                  <div class="custom-menu-name">
                    <a-form-item label="名称" field="name" required>
                      <a-input
                        v-model="formData.name"
                        placeholder="仅支持中文或数字"
                      />
                      <template #extra>
                        <div
                          >仅支持中英文和数字，字数不超过4个汉字或8个字母。
                        </div>
                      </template>
                      <!--                        <p class="input-tip">仅支持中英文和数字，字数不超过4个汉字或8个字母。</p>-->
                    </a-form-item>
                  </div>
                  <div class="custom-menu-type">
                    <a-form-item label="消息类型">
                      <a-radio-group
                        v-model="menuType"
                        :default-value="menuType"
                        :options="menuTypeOptions"
                        class="ml-4"
                      />
                    </a-form-item>
                  </div>
                  <div v-if="menuType === 'click'" class="custom-menu-content">
                    <a-form-item label="菜单内容" field="key" required>
                      <a-input
                        v-model="formData.key"
                        placeholder="仅支持中文或数字"
                      />
                      <template #extra>
                        <div>key值为管理平台创建好的功能key</div>
                      </template>
                    </a-form-item>
                  </div>
                  <div v-if="menuType === 'view'" class="custom-menu-content">
                    <a-form-item label="网页链接" field="url" required>
                      <a-input
                        v-model="formData.url"
                        placeholder="公众号链接"
                      />
                      <template #extra>
                        <div>跳转连接推荐使用安全域名https://</div>
                      </template>
                    </a-form-item>
                  </div>
                  <div
                    v-if="menuType === 'miniprogram'"
                    class="custom-menu-content"
                  >
                    <a-form-item label="Appid" field="appid" required>
                      <a-input
                        v-model="formData.appid"
                        placeholder="小程序Appid"
                      />
                      <template #extra>
                        <div
                          >输入对应的小程序Appid 示例：wxd027d2b162044fd5
                        </div>
                      </template>
                    </a-form-item>
                    <a-form-item label="页面路径" field="pagepath" required>
                      <a-input
                        v-model="formData.pagepath"
                        placeholder="页面路径"
                      />
                      <template #extra>
                        <div
                          >输入对应的小程序页面路径 示例：/pages/index/index
                        </div>
                      </template>
                    </a-form-item>
                  </div>
                </a-form>
                <div class="custom-menu-delete">
                  <div class="input-title">
                    <div class="delete-btn">删除菜单</div>
                  </div>
                </div>
              </div>
              <!--              有二级菜单-->
              <div
                v-if="
                  isHaveSubMenu() && selectMenuId !== 0 && selectMenuSubId === 0
                "
                class="custom-have-menu-box"
              >
                <h3 style="padding-bottom: 25px">菜单信息</h3>
                <a-form ref="formVisible" :model="formData" :rules="rule">
                  <div class="custom-menu-name">
                    <a-form-item label="名称" field="name" required>
                      <a-input
                        v-model="formData.name"
                        placeholder="仅支持中文或数字"
                      />
                      <template #extra>
                        <div
                          >仅支持中英文和数字，字数不超过4个汉字或8个字母。
                        </div>
                      </template>
                      <!--                        <p class="input-tip">仅支持中英文和数字，字数不超过4个汉字或8个字母。</p>-->
                    </a-form-item>
                  </div>
                  <div class="custom-menu-delete">
                    <div class="input-title">
                      <a-button type="text" class="delete-btn"
                        >删除菜单
                      </a-button>
                    </div>
                  </div>
                </a-form>
              </div>
            </div>
            <!--            没有菜单-->
            <div v-if="menuData.button.length === 0" class="attribute-box">
              <a-empty
                description="你未添加自定义菜单，点击左侧添加菜单为公众号创建菜单栏。"
              />
            </div>
            <!--            右侧底部按钮组-->
            <div class="menu__edit-action-bar">
              <div class="menu__edit-action-inner">
                <div class="menu__edit-button">
                  <a-button @click="preview">预览</a-button>
                  <a-button type="primary" @click="saveMenu">保存</a-button>
                  <a-button type="primary" @click="submitMenuData"
                    >保存并发布
                  </a-button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </a-col>
    </a-row>
  </div>
</template>

<script setup lang="ts">
  import { onMounted, ref } from 'vue';
  import { Message, Modal } from '@arco-design/web-vue';
  import { IconPlus } from '@arco-design/web-vue/es/icon';
  import {
    getOAMenuTree,
    GetOAMenuTreeReply,
    syncOAMenu,
  } from '@/api/we-public-menu';

  const selectMenuId = ref(0);
  const selectMenuSubId = ref(0);

  // 菜单类型
  const menuType = ref('');
  const menuTypeOptions = ref([
    { label: '推送事件消息', value: 'click' },
    { label: '网页跳转', value: 'view' },
    { label: '小程序跳转', value: 'miniprogram' },
  ]);
  // 表单数据
  const formData = ref<Record<string, any>>({});

  // 验证规则
  const rule = ref({
    name: [
      { required: true, message: '请输入菜单名称', trigger: 'blur' },
      { minLength: 2, message: '请输入菜单名称 最低5位字符', trigger: 'blur' },
      { maxLength: 4, message: '请输入菜单名称 最多4位字符', trigger: 'blur' },
    ],
    url: [
      { required: true, message: '请输入网页链接', trigger: 'blur' },
      {
        match: /^(http|https):\/\//,
        message: '请输入正确的网页链接',
        trigger: 'blur',
      },
    ],
    appid: [
      { required: true, message: '请输入小程序Appid', trigger: 'blur' },
      {
        match: /^[a-zA-Z0-9]{18}$/,
        message: '请输入正确的小程序Appid',
        trigger: 'blur',
      },
    ],
    pagepath: [
      { required: true, message: '请输入页面路径', trigger: 'blur' },
      // 是否是小程序页面路径/pages/index/index
      {
        match: /^\/pages\/[a-zA-Z0-9]+\/[a-zA-Z0-9]+$/,
        message: '请输入正确的页面路径',
        trigger: 'blur',
      },
    ],
    key: [{ required: true, message: '请输入菜单内容' }],
  });

  // 菜单数据
  const menuData = ref({
    button: [] as any[],
    matchrule: {
      tag_id: '2',
      sex: '1',
      country: '中国',
      province: '广东',
      city: '广州',
      // client_platform_type: '2',
      language: 'zh_CN',
    },
  });

  // 保存菜单数据
  const saveMenu = async () => {
    // 如果没有选择菜单，不保存
    if (selectMenuId.value === 0) {
      return;
    }
    // 找到对应的菜单
    const menu = menuData.value.button.find(
      (item: any) => item.id === selectMenuId.value
    );
    // 如果没有选择二级菜单
    if (selectMenuSubId.value === 0) {
      // 如果没有二级菜单，修改完整的一级菜单 如果有二级菜单，修改一级菜单的name
      if (menu.subButton.length === 0) {
        menu.name = formData.value.name;
        menu.type = menuType.value;
        if (menuType.value === 'click') {
          menu.key = formData.value.key;
          // 把当前一级菜单除了name、subButton、id、type、key字段外 删除一级菜单的其他字段
          const { name, subButton, id, type, key } = menu;
          menuData.value.button.splice(
            menuData.value.button.findIndex(
              (item: any) => item.id === selectMenuId.value
            ),
            1,
            {
              name,
              subButton,
              id,
              type,
              key,
            }
          );
        } else if (menuType.value === 'view') {
          menu.url = formData.value.url;
          // 把当前一级菜单除了name、subButton、id、type、url字段外 删除一级菜单的其他字段
          const {
            name: name1,
            subButton: subButton1,
            id: id1,
            type: type1,
            url,
          } = menu;
          menuData.value.button.splice(
            menuData.value.button.findIndex(
              (item: any) => item.id === selectMenuId.value
            ),
            1,
            {
              name: name1,
              subButton: subButton1,
              id: id1,
              type: type1,
              url,
            }
          );
        } else if (menuType.value === 'miniprogram') {
          menu.appid = formData.value.appid;
          menu.pagepath = formData.value.pagepath;
          // 把当前一级菜单除了name、subButton、id、type、appid、pagepath字段外 删除一级菜单的其他字段
          const {
            name: name2,
            subButton: subButton2,
            id: id2,
            type: type2,
            appid,
            pagepath,
          } = menu;
          menuData.value.button.splice(
            menuData.value.button.findIndex(
              (item: any) => item.id === selectMenuId.value
            ),
            1,
            {
              name: name2,
              subButton: subButton2,
              id: id2,
              type: type2,
              appid,
              pagepath,
            }
          );
        }
      } else {
        menu.name = formData.value.name;
      }
    } else {
      const subMenu = menu.subButton.find(
        (item: any) => item.id === selectMenuSubId.value
      );
      subMenu.name = formData.value.name;
      subMenu.type = menuType.value;
      if (menuType.value === 'click') {
        subMenu.key = formData.value.key;
        // 把当前二级菜单除了name、id、type、key字段外 删除一级菜单的其他字段
        const { name, id, type, key } = subMenu;
        menu.subButton.splice(
          menu.subButton.findIndex(
            (item: any) => item.id === selectMenuSubId.value
          ),
          1,
          {
            name,
            id,
            type,
            key,
          }
        );
      } else if (menuType.value === 'view') {
        subMenu.url = formData.value.url;
        // 把当前二级菜单除了name、id、type、url字段外 删除一级菜单的其他字段
        const { name: name1, id: id1, type: type1, url } = subMenu;
        menu.subButton.splice(
          menu.subButton.findIndex(
            (item: any) => item.id === selectMenuSubId.value
          ),
          1,
          {
            name: name1,
            id: id1,
            type: type1,
            url,
          }
        );
      } else if (menuType.value === 'miniprogram') {
        subMenu.appid = formData.value.appid;
        subMenu.pagepath = formData.value.pagepath;
        // 把当前二级菜单除了name、id、type、appid、pagepath字段外 删除一级菜单的其他字段
        const { name: name2, id: id2, type: type2, appid, pagepath } = subMenu;
        menu.subButton.splice(
          menu.subButton.findIndex(
            (item: any) => item.id === selectMenuSubId.value
          ),
          1,
          {
            name: name2,
            id: id2,
            type: type2,
            appid,
            pagepath,
          }
        );
      }
    }
  };

  // 提交菜单数据
  const submitMenuData = async () => {
    if (menuData.value.button.length === 0) {
      Modal.info({
        title: '提醒',
        content: '请先添加菜单',
      });
      return;
    }
    await saveMenu();
    const res = await syncOAMenu(menuData.value);
    if (res.data.success) {
      Message.info('同步菜单成功');
    } else {
      Message.error('同步菜单失败');
    }
  };

  // 预览
  const preview = () => {
    Modal.info({
      title: '提醒',
      content: '开源版不支持预览\n如有需要专业版本请联系作者',
    });
  };

  // 是否包含二级菜单
  const isHaveSubMenu = () => {
    if (menuData.value.button.length > 0) {
      const menuId = selectMenuId.value;
      const menu = menuData.value.button.find(
        (item: any) => item.id === menuId
      );
      // 如果存在二级菜单
      return menu.subButton && menu.subButton.length > 0;
    }
    return false;
  };

  // 选择菜单
  const selectMenu = async (item: number) => {
    formData.value = {};
    selectMenuId.value = item;
    selectMenuSubId.value = 0;
    const menu = menuData.value.button.find(
      (mItem: any) => mItem.id === selectMenuId.value
    );
    // 如果没有二级菜单，将form表单的数据设置为一级菜单的数据
    if (
      menuData.value.button.find(
        (mItem: any) => mItem.id === selectMenuId.value
      ).subButton.length === 0
    ) {
      formData.value.name = menu.name;
      menuType.value = menu.type;
      switch (menuType.value) {
        case 'click':
          formData.value.key = menu.key;
          break;
        case 'view':
          formData.value.url = menu.url;
          break;
        case 'miniprogram':
          formData.value.appid = menu.appid;
          formData.value.pagepath = menu.pagepath;
          break;
        default:
          break;
      }
    } else {
      formData.value.name = menu.name;
    }
  };

  // 添加一级菜单
  const addMainMenu = async () => {
    formData.value = {};
    if (menuData.value.button.length < 3) {
      let menuId = 0;
      if (menuData.value.button.length === 0) {
        menuId = 1;
      } else {
        menuId = menuData.value.button[menuData.value.button.length - 1].id + 1;
      }
      menuData.value.button.push({
        id: menuId,
        name: '菜单名称',
        type: 'click',
        key: '',
        subButton: [],
      });
      // 如果新增菜单将菜单id设置为新增菜单的id
      selectMenuId.value = menuId;
      await selectMenu(selectMenuId.value);
    } else {
      Message.warning('最多只能添加三个一级菜单');
    }
  };

  // 选择二级菜单
  const selectSubMenu = async (item: number) => {
    formData.value = {};
    selectMenuSubId.value = item;
    const menu = menuData.value.button.find(
      (mItem: any) => mItem.id === selectMenuId.value
    );
    const subMenu = menu.subButton.find(
      (mItem: any) => mItem.id === selectMenuSubId.value
    );
    formData.value.name = subMenu.name;
    switch (subMenu.type) {
      case 'click':
        menuType.value = 'click';
        formData.value.key = subMenu.key;
        break;
      case 'view':
        menuType.value = 'view';
        formData.value.url = subMenu.url;
        break;
      case 'miniprogram':
        menuType.value = 'miniprogram';
        formData.value.appid = subMenu.appid;
        formData.value.pagepath = subMenu.pagepath;
        break;
      default:
        break;
    }
  };

  // 添加二级菜单
  const addSubMenu = async () => {
    formData.value = {};
    const menuId = selectMenuId.value;
    const menu = menuData.value.button.find((item: any) => item.id === menuId);
    // 如果二级菜单小于5个
    if (menu.subButton.length < 5) {
      let subMenuId = 0;
      if (menu.subButton.length === 0) {
        // 如果没有二级菜单，设置id为1
        subMenuId = 1;
      } else {
        // 如果有二级菜单，设置id为最后一个二级菜单的id+1
        subMenuId = menu.subButton[menu.subButton.length - 1].id + 1;
      }
      menu.subButton.push({
        // 添加二级菜单
        id: menu.subButton.length + 1,
        name: '子菜单名称',
        type: 'click',
        key: '',
      });
      selectMenuSubId.value = subMenuId;
      // 将选中二级菜单id设置为新增二级菜单的id
      await selectSubMenu(selectMenuSubId.value);
      // 把当前一级菜单除了name和subButton和id字段外 删除一级菜单的其他字段
      const { name, subButton, id } = menu;
      menuData.value.button.splice(
        menuData.value.button.findIndex((item: any) => item.id === menuId),
        1,
        {
          name,
          subButton,
          id,
        }
      );
    } else {
      Message.warning('最多只能添加五个二级菜单');
    }
  };

  // 往左移动选择菜单
  const onLeftMove = async () => {
    let previousId = null;
    for (let i = 0; i < menuData.value.button.length; i += 1) {
      if (menuData.value.button[i].id === selectMenuId.value) {
        if (i > 0) {
          previousId = menuData.value.button[i - 1].id;
        }
        break; // 找到目标后可以提前退出循环
      }
    }
    await selectMenu(previousId);
  };

  // 删除对应id的菜单
  const onDelete = async () => {
    const menuId = selectMenuId.value;
    for (let i = 0; i < menuData.value.button.length; i += 1) {
      if (menuData.value.button[i].id === menuId) {
        menuData.value.button.splice(i, 1);
        break; // 可以提高性能，因为一旦找到要删除的对象就不需要再继续循环
      }
    }
  };
  // 删除对应id的二级菜单
  const onDeleteSub = async () => {
    const menuId = selectMenuId.value;
    const subMenuId = selectMenuSubId.value;
    for (let i = 0; i < menuData.value.button.length; i += 1) {
      if (menuData.value.button[i].id === menuId) {
        for (let j = 0; j < menuData.value.button[i].subButton.length; j += 1) {
          if (menuData.value.button[i].subButton[j].id === subMenuId) {
            menuData.value.button[i].subButton.splice(j, 1);

            // 如果删除的是最后一个二级菜单，将一级菜单的id设置为选中
            if (menuData.value.button[i].subButton.length === 0) {
              selectMenuId.value = menuId;
              selectMenu(selectMenuId.value);
            } else if (j === menuData.value.button[i].subButton.length) {
              selectMenuSubId.value =
                menuData.value.button[i].subButton[j - 1].id;
              selectSubMenu(selectMenuSubId.value);
            } else {
              selectMenuSubId.value = menuData.value.button[i].subButton[j].id;
              selectSubMenu(selectMenuSubId.value);
            }

            break; // 可以提高性能，因为一旦找到要删除的对象就不需要再继续循环
          }
        }
      }
    }
  };

  // 往右移动选择菜单
  const onRightMove = async () => {
    let nextId = null;

    for (let i = 0; i < menuData.value.button.length; i += 1) {
      if (menuData.value.button[i].id === selectMenuId.value) {
        if (i < menuData.value.button.length - 1) {
          nextId = menuData.value.button[i + 1].id;
        }
        break; // 找到目标后可以提前退出循环
      }
    }
    await selectMenu(nextId);
  };

  const menuDataList = ref<GetOAMenuTreeReply>();

  onMounted(async () => {
    const res = await getOAMenuTree();
    menuDataList.value = res.data;
  });
</script>

<style scoped lang="scss">
  body {
    overflow: hidden;
  }

  .header-box {
    display: flex;
    -webkit-box-align: center;
    flex: 1;
    height: 80px;
    padding-bottom: 10px;

    .info_tips_wrap {
      display: flex;
      align-items: center;
      justify-content: center;
      background-color: #fff;
      border-radius: 8px 0 0 8px;
      padding-left: 18px;

      .suc_icon {
        width: 24px;
        height: 24px;
        background: url(data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjQiIGhlaWdodD0iMjQiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PHBhdGggZmlsbC1ydWxlPSJldmVub2RkIiBjbGlwLXJ1bGU9ImV2ZW5vZGQiIGQ9Ik0xMS44NzcgMkMxNy41MjMgMiAyMiA2LjQ3NyAyMiAxMi4xMjMgMjIgMTcuNTIzIDE3LjUyMyAyMiAxMS44NzcgMjIgNi40NzYgMjIgMiAxNy41MjMgMiAxMi4xMjMgMiA2LjQ3OCA2LjQ3NyAyIDExLjg3NyAyem0tNC44MSA5LjkzMmEuMzM5LjMzOSAwIDAwLjAyMy40MzJsMy4wNDIgMy4xMzdjLjExLjExMy4yODIuMTE1LjM4OS4wMWw2Ljg4LTYuNzQ5YS4yNzIuMjcyIDAgMDAtLjAwMS0uMzgzbC0uMTg3LS4xODRhLjMxLjMxIDAgMDAtLjQwNC0uMDIxbC02LjI3NSA1LjIxYS4zNjUuMzY1IDAgMDEtLjQzMS4wMDFMNy42OTkgMTEuNTVhLjI3NS4yNzUgMCAwMC0uMzk1LjA2bC0uMjM3LjMyM3oiIGZpbGw9IiMwN0MxNjAiLz48L3N2Zz4=)
          50% no-repeat;
      }
    }

    .content {
      display: inline-flex;
      float: left;
      align-items: center;
      background-color: #fff;
      border-radius: 0 8px 8px 0;
      flex: 1;
      height: 100%;
      // 不换行
      white-space: nowrap;

      .in-content {
        padding-left: 15px;

        .title {
          font-size: 13px;
        }
      }
    }
  }

  .container-box {
    display: flex;
    justify-content: center;
    max-height: 634px;

    .phone-box {
      position: relative;
      display: block;
      width: 300px;
      min-width: 300px;
      min-height: 550px;
      max-height: 634px;
      background: #ffffff;
      border-radius: 8px;
      box-shadow: 0 1px 5px 0 rgba(0, 0, 0, 0.05);

      .menu__preview-hd {
        background-image: url(data:image/png;base64,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);
        background-size: 100% 100%;
        height: 70px;
      }

      .menu__preview-bd {
        background: #fdfdfd;
        min-height: 480px;
        max-height: 564px;
        border-radius: 0 0 8px 8px;

        .menu__preview-bottom {
          position: absolute;
          left: 0;
          bottom: 0;
          width: 100%;
          height: 48px;
          display: -webkit-box;
          display: -webkit-flex;
          display: -ms-flexbox;
          display: flex;
          -webkit-box-pack: justify;
          -webkit-justify-content: space-between;
          -ms-flex-pack: justify;
          justify-content: space-between;
          -webkit-box-shadow: inset 0 0.5px 0 #e3e4e5;
          box-shadow: inset 0 0.5px 0 #e3e4e5;
          -webkit-box-align: center;
          -webkit-align-items: center;
          -ms-flex-align: center;
          align-items: center;
          background-color: #fff;
          border-radius: 0 0 8px 8px;

          .menu__keyboard {
            display: inline-flex;
            align-items: center;
            //width: auto;
            height: 100%;

            .menu__keyboard-icon {
              display: inline-flex;
              align-items: center;
              width: 40px;
              height: 100%;
              background: url(data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNDAiIGhlaWdodD0iNjgiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PHBhdGggZmlsbC1ydWxlPSJldmVub2RkIiBjbGlwLXJ1bGU9ImV2ZW5vZGQiIGQ9Ik0xNy40MjkgMjAuNjEyTDIwIDE4bDIuNTcyIDIuNjEyaC01LjE0M3pNMjAgNDMuOTY5Yy01LjUyMyAwLTEwLTQuNTQ4LTEwLTEwLjE1NyAwLTUuNjEgNC40NzctMTAuMTU3IDEwLTEwLjE1N3MxMCA0LjU0OCAxMCAxMC4xNTdjMCA1LjYxLTQuNDc3IDEwLjE1Ny0xMCAxMC4xNTd6bTAtMS4yMmM0Ljg2IDAgOC44LTQgOC44LTguOTM3IDAtNC45MzYtMy45NC04LjkzOC04LjgtOC45MzgtNC44NiAwLTguOCA0LjAwMi04LjggOC45MzggMCA0LjkzNiAzLjk0IDguOTM4IDguOCA4LjkzOHptLTMuNS0xM2gtMnYyLjAzMmgydi0yLjAzMnptMSAwaDJ2Mi4wMzJoLTJ2LTIuMDMyem01IDBoLTJ2Mi4wMzJoMnYtMi4wMzJ6bTEgMGgydjIuMDMyaC0ydi0yLjAzMnptLTcgMy4wNDdoLTJ2Mi4wMzJoMnYtMi4wMzJ6bTEgMGgydjIuMDMyaC0ydi0yLjAzMnpNMjMgMzguODlWMzYuODZoLTZ2Mi4wMzFoNnptLTIuNS02LjA5NGgydjIuMDMyaC0ydi0yLjAzMnptNSAwaC0ydjIuMDMyaDJ2LTIuMDMyeiIgZmlsbD0iI0IxQjJCMyIvPjwvc3ZnPg==)
                50% no-repeat;
              //background-size: cover;
            }
          }

          .menu-list {
            display: inline-flex;
            align-items: center;
            flex: 1;
            //margin-right: 0.5%;
            position: relative;

            .menu-add-no-menu {
              //position: absolute;
              display: inline-flex;
              justify-content: center;
              align-items: center;
              width: 100%;
              height: 48px;
              flex: 1;
              -webkit-box-flex: 1;
              border: 2px dashed rgba(7, 193, 96, 0.6);
              border-radius: 3px;
              //padding-right: 1vw;

              .add-no-menu {
                height: 100%;
                width: 100%;
                display: flex;
                justify-content: center;
                align-items: center;
              }

              .menu__add-icon {
                vertical-align: middle;
                display: inline-block;
                background: url(data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMTYiIGhlaWdodD0iMTYiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PHBhdGggZmlsbC1ydWxlPSJldmVub2RkIiBjbGlwLXJ1bGU9ImV2ZW5vZGQiIGQ9Ik03LjUgMTMuNWEuNS41IDAgMDAxIDB2LTVoNWEuNS41IDAgMDAwLTFoLTV2LTVhLjUuNSAwIDAwLTEgMHY1aC01YS41LjUgMCAwMDAgMWg1djV6IiBmaWxsPSIjMDdDMTYwIi8+PC9zdmc+)
                  no-repeat;
                width: 20px;
                height: 17px;
              }

              .add-no-menu span {
                color: #07c160;
              }
            }

            .add-button {
              display: flex;
              justify-content: center;
              align-items: center;
              width: 100%;
              line-height: 48px;
              flex: 1;
              border-radius: 3px;
              //padding-right: 1vw;
              position: relative;

              .menu_button_preview-line {
                position: absolute;
                width: 0.5px;
                height: 33px;
                background: #e3e4e5;
                left: 0;
                bottom: 15px;
                right: 0;
              }

              .a-button {
                height: 48px;
                width: 100%;
                display: inline-flex;
                justify-content: center;
                align-items: center;

                .icon14_menu_add {
                  background: url(data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjAiIGhlaWdodD0iMjAiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PHBhdGggZmlsbC1ydWxlPSJldmVub2RkIiBjbGlwLXJ1bGU9ImV2ZW5vZGQiIGQ9Ik05LjM3NSAxN2EuNS41IDAgMDAuNS41aC4yNWEuNS41IDAgMDAuNS0uNXYtNi4zNzVIMTdhLjUuNSAwIDAwLjUtLjV2LS4yNWEuNS41IDAgMDAtLjUtLjVoLTYuMzc1VjNhLjUuNSAwIDAwLS41LS41aC0uMjVhLjUuNSAwIDAwLS41LjV2Ni4zNzVIM2EuNS41IDAgMDAtLjUuNXYuMjVhLjUuNSAwIDAwLjUuNWg2LjM3NVYxN3oiIGZpbGw9IiM0QzRENEUiLz48L3N2Zz4=)
                    no-repeat;
                  width: 19px;
                  height: 22px;
                }
              }
            }

            .menu-list-item {
              position: relative;
              display: flex;
              justify-content: center;
              align-items: center;
              line-height: 48px;
              flex: 1;
              border-radius: 2px;

              .menu__preview-line {
                position: absolute;
                width: 0.5px;
                height: 33px;
                background: #e3e4e5;
                left: 0;
                bottom: 15px;
                right: 0;
              }

              .menu-item {
                height: 100%;
                flex: 1;
                display: flex;
                justify-content: center;
                align-items: center;
                cursor: pointer;
              }

              .submenu {
                position: absolute;
                bottom: 62px;
                left: 0;
                width: 100%;
                border: 1px solid #e3e4e5;
                border-top: 0 solid #e3e4e5;
                background: #fff;
                border-radius: 4px;

                .menu-item-sub {
                  line-height: 48px;
                  width: auto;
                  //color: #616161;
                  display: flex;
                  justify-content: center;
                  align-items: center;
                  cursor: pointer;
                  position: relative;

                  .sub-menu-bar {
                    position: absolute;
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    right: -50px;
                    background: hsla(0, 0%, 100%, 0.8);
                    box-shadow: 0 1px 6px #e4e8eb;
                    border-radius: 24px;

                    .sub-menu-bar-box {
                      display: flex;
                      background: #fff;
                      border-radius: 24px;

                      .sub-center-bar {
                        width: 38px;
                        height: 38px;
                        background: url(data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjQiIGhlaWdodD0iMjQiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PHBhdGggZmlsbC1ydWxlPSJldmVub2RkIiBjbGlwLXJ1bGU9ImV2ZW5vZGQiIGQ9Ik0xNC4yOTIgNC4yMDhhLjQ1OS40NTkgMCAwMC0uNDU5LS40NThoLTMuNjY2YS40NTkuNDU5IDAgMDAtLjQ1OS40NTh2LjY0Mmg0LjU4NHYtLjY0MnptMy42MDEgMi42NTlsLS43NDkgMTIuNTc1YTEuODMzIDEuODMzIDAgMDEtMS44MyAxLjcyNUg4LjY4NmMtLjk3IDAtMS43NzMtLjc1Ni0xLjgzLTEuNzI1TDYuMTA3IDYuODY3SDQuMjA4di0uNjQyYzAtLjI1My4yMDYtLjQ1OC40NTktLjQ1OGgxNC42NjZjLjI1MyAwIC40NTkuMjA1LjQ1OS40NTh2LjY0MmgtMS44OTl6bS0xMC42ODMgMGwuNzQ0IDEyLjUxYy4wMjMuMzg3LjM0NC42OS43MzIuNjloNi42MjhjLjM4OCAwIC43MS0uMzAyLjczMi0uNjlsLjc0NS0xMi41MUg3LjIxek05LjcwOCA5LjI1aDEuMWwuNDU5IDguMjVoLTEuMWwtLjQ1OS04LjI1em00LjU4NCAwaC0xLjFsLS40NTkgOC4yNWgxLjFsLjQ1OS04LjI1eiIgZmlsbD0iIzRBNEE1MSIvPjxwYXRoIGQ9Ik05LjcwOCA0Ljg1aC0uMXYuMWguMXYtLjF6bTQuNTg0IDB2LjFoLjF2LS4xaC0uMXptMi44NTIgMTQuNTkybC0uMS0uMDA2LjEuMDA2em0uNzQ5LTEyLjU3NXYtLjFoLS4wOTRsLS4wMDYuMDk0LjEuMDA2em0tMi41NzkgMTQuM3YtLjEuMXptLTguNDU4LTEuNzI1bC4xLS4wMDYtLjEuMDA2ek02LjEwNyA2Ljg2N2wuMS0uMDA2LS4wMDYtLjA5NGgtLjA5NHYuMXptLTEuODk5IDBoLS4xdi4xaC4xdi0uMXptMTUuNTg0IDB2LjFoLjF2LS4xaC0uMXpNNy45NTQgMTkuMzc3bC4xLS4wMDYtLjEuMDA2ek03LjIxIDYuODY3di0uMWgtLjEwNmwuMDA2LjEwNi4xLS4wMDZ6bTguMTA0IDEzLjJ2LS4xLjF6bS43MzItLjY5bC0uMS0uMDA2LjEuMDA2em0uNzQ1LTEyLjUxbC4xLjAwNi4wMDYtLjEwNmgtLjEwNnYuMXpNMTAuODA4IDkuMjVsLjEtLjAwNi0uMDA1LS4wOTRoLS4wOTV2LjF6bS0xLjEgMHYtLjFoLS4xMDZsLjAwNi4xMDYuMS0uMDA2em0xLjU1OSA4LjI1di4xaC4xMDZsLS4wMDYtLjEwNi0uMS4wMDZ6bS0xLjEgMGwtLjEuMDA2LjAwNS4wOTRoLjA5NXYtLjF6bTMuMDI1LTguMjV2LS4xaC0uMDk1bC0uMDA1LjA5NC4xLjAwNnptMS4xIDBsLjEuMDA2LjAwNi0uMTA2aC0uMTA2di4xem0tMS41NTkgOC4yNWwtLjEtLjAwNi0uMDA2LjEwNmguMTA2di0uMXptMS4xIDB2LjFoLjA5NWwuMDA1LS4wOTQtLjEtLjAwNnptMC0xMy42NWMuMTk4IDAgLjM1OS4xNi4zNTkuMzU4aC4yYS41NTkuNTU5IDAgMDAtLjU1OS0uNTU4di4yem0tMy42NjYgMGgzLjY2NnYtLjJoLTMuNjY2di4yem0tLjM1OS4zNThjMC0uMTk3LjE2LS4zNTguMzU5LS4zNTh2LS4yYS41NTkuNTU5IDAgMDAtLjU1OS41NThoLjJ6bTAgLjY0MnYtLjY0MmgtLjJ2LjY0MmguMnptNC40ODQtLjFIOS43MDh2LjJoNC41ODR2LS4yem0tLjEtLjU0MnYuNjQyaC4ydi0uNjQyaC0uMnptMy4wNTIgMTUuMjRsLjc0OS0xMi41NzUtLjItLjAxMi0uNzQ5IDEyLjU3NS4yLjAxMnptLTEuOTMgMS44MTljMS4wMjMgMCAxLjg3LS43OTcgMS45My0xLjgybC0uMi0uMDFhMS43MzMgMS43MzMgMCAwMS0xLjczIDEuNjN2LjJ6bS02LjYyOCAwaDYuNjI4di0uMkg4LjY4NnYuMnptLTEuOTMtMS44MmExLjkzMyAxLjkzMyAwIDAwMS45MyAxLjgydi0uMmExLjczMyAxLjczMyAwIDAxLTEuNzMtMS42M2wtLjIuMDF6TTYuMDA3IDYuODc0bC43NSAxMi41NzUuMTk5LS4wMTItLjc1LTEyLjU3NS0uMTk5LjAxMnptLTEuNzk5LjA5NGgxLjg5OXYtLjJINC4yMDh2LjJ6bS0uMS0uNzQydi42NDJoLjJ2LS42NDJoLS4yem0uNTU5LS41NThhLjU1OS41NTkgMCAwMC0uNTU5LjU1OGguMmMwLS4xOTcuMTYxLS4zNTguMzU5LS4zNTh2LS4yem0xNC42NjYgMEg0LjY2N3YuMmgxNC42NjZ2LS4yem0uNTU5LjU1OGEuNTU5LjU1OSAwIDAwLS41NTktLjU1OHYuMmMuMTk4IDAgLjM1OS4xNi4zNTkuMzU4aC4yem0wIC42NDJ2LS42NDJoLS4ydi42NDJoLjJ6bS0xLjk5OS4xaDEuODk5di0uMmgtMS44OTl2LjJ6bS05Ljg0IDEyLjQwNEw3LjMxIDYuODYxbC0uMi4wMTIuNzQ0IDEyLjUxLjItLjAxMnptLjYzMy41OTZhLjYzNC42MzQgMCAwMS0uNjMyLS41OTZsLS4yLjAxMmMuMDI2LjQ0LjM5MS43ODQuODMyLjc4NHYtLjJ6bTYuNjI4IDBIOC42ODZ2LjJoNi42Mjh2LS4yem0uNjMyLS41OTZhLjYzMy42MzMgMCAwMS0uNjMyLjU5NnYuMmEuODMzLjgzMyAwIDAwLjgzMi0uNzg0bC0uMi0uMDEyem0uNzQ1LTEyLjUxbC0uNzQ1IDEyLjUxLjIuMDEyLjc0NS0xMi41MS0uMi0uMDEyem0tOS40ODEuMTA2aDkuNTgxdi0uMkg3LjIxdi4yem0zLjU5OCAyLjE4M2gtMS4xdi4yaDEuMXYtLjJ6bS41NTkgOC4zNDRsLS40Ni04LjI1LS4xOTkuMDEyLjQ2IDguMjUuMTk5LS4wMTJ6bS0xLjIuMTA2aDEuMXYtLjJoLTEuMXYuMnptLS41NTktOC4zNDRsLjQ2IDguMjUuMTk5LS4wMTItLjQ2LTguMjUtLjE5OS4wMTJ6bTMuNTg0LjA5NGgxLjF2LS4yaC0xLjF2LjJ6bS0uMzYgOC4xNTZsLjQ2LTguMjUtLjItLjAxMi0uNDU5IDguMjUuMi4wMTJ6bTEuMDAxLS4xMDZoLTEuMXYuMmgxLjF2LS4yem0uMzYtOC4xNTZsLS40NiA4LjI1LjIuMDEyLjQ1OS04LjI1LS4yLS4wMTJ6IiBmaWxsPSIjNEE0QTUxIi8+PC9zdmc+)
                          50% no-repeat;
                      }

                      .bar-hover:hover {
                        background-color: #e3e3e4;
                        border-radius: 24px;
                        cursor: pointer;
                      }
                    }
                  }
                }

                .menu-item-border-color {
                  border-top: 1px solid #e7e7eb;
                }

                .add-button-sub {
                  display: flex;
                  justify-content: center;
                  align-items: center;
                  border-top: 1px solid #e7e7eb;
                  height: 44px;

                  .arco-btn {
                    width: 100%;
                    height: 100%;
                    border-radius: 0;
                    border: 0;
                    background: #fff;
                    color: #07c160;
                    font-size: 14px;
                    font-weight: 500;
                    cursor: pointer;

                    &:hover {
                      background: #fff;
                      color: #07c160;
                    }
                  }
                }

                .arrow {
                  position: absolute;
                  left: 50%;
                  margin-left: -6px;
                }

                .arrow_out {
                  bottom: -6px;
                  display: inline-block;
                  width: 0;
                  height: 0;
                  border-width: 6px;
                  border-style: dashed;
                  border-color: transparent;
                  border-bottom-width: 0;
                  border-top-color: #d0d0d0;
                  border-top-style: solid;
                }

                .arrow_in {
                  bottom: -5px;
                  display: inline-block;
                  width: 0;
                  height: 0;
                  border-width: 6px;
                  border-style: dashed;
                  border-color: transparent;
                  border-bottom-width: 0;
                  border-top-color: #fafafa;
                  border-top-style: solid;
                }

                .arrow_out_select {
                  border-top-color: #07c160 !important;
                }

                .arrow_in_select {
                  z-index: 1 !important;
                  bottom: -4px !important;
                }
              }
            }

            .menu-box-shadow {
              border: 1.5px solid #07c160;
              background-color: rgba(7, 193, 96, 0.05);
              height: 48px;
            }

            .menu-item-color {
              color: #07c160;
            }

            .menu-bar {
              position: absolute;
              display: flex;
              align-items: center;
              justify-content: center;
              width: 80%;
              height: 80px;
              top: 40px;
              left: 50%;
              transform: translateX(-50%);

              .menu-bar-box {
                display: flex;
                background: #fff;
                border-radius: 24px;

                .left-bar {
                  width: 38px;
                  height: 38px;
                  transform: rotate(-90deg);
                  background: url(data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjQiIGhlaWdodD0iMjQiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PHBhdGggZD0iTTEwLjg4NiA0Ljc2NHYxNS43OGMwIC4xMjIuMDk4LjIyLjIyLjIyaDEuMDlhLjIyLjIyIDAgMDAuMjE4LS4yMlY0Ljc2NWgtMS41Mjh6IiBmaWxsPSIjNEE0QTUxIi8+PHBhdGggZD0iTTExLjk0NCAzLjE2M2w1Ljg5IDYuNDVhLjUuNSAwIDAxLS42NS43NUwxMS44NTcgNi43NGEuNS41IDAgMDAtLjU2MiAwbC01LjMzIDMuNjIzYS41LjUgMCAwMS0uNjUtLjc1bDUuODktNi40NWEuNS41IDAgMDEuNzQgMHoiIGZpbGw9IiM0QTRBNTEiLz48L3N2Zz4=)
                    50% no-repeat;
                }

                .center-bar {
                  width: 38px;
                  height: 38px;
                  background: url(data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjQiIGhlaWdodD0iMjQiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PHBhdGggZmlsbC1ydWxlPSJldmVub2RkIiBjbGlwLXJ1bGU9ImV2ZW5vZGQiIGQ9Ik0xNC4yOTIgNC4yMDhhLjQ1OS40NTkgMCAwMC0uNDU5LS40NThoLTMuNjY2YS40NTkuNDU5IDAgMDAtLjQ1OS40NTh2LjY0Mmg0LjU4NHYtLjY0MnptMy42MDEgMi42NTlsLS43NDkgMTIuNTc1YTEuODMzIDEuODMzIDAgMDEtMS44MyAxLjcyNUg4LjY4NmMtLjk3IDAtMS43NzMtLjc1Ni0xLjgzLTEuNzI1TDYuMTA3IDYuODY3SDQuMjA4di0uNjQyYzAtLjI1My4yMDYtLjQ1OC40NTktLjQ1OGgxNC42NjZjLjI1MyAwIC40NTkuMjA1LjQ1OS40NTh2LjY0MmgtMS44OTl6bS0xMC42ODMgMGwuNzQ0IDEyLjUxYy4wMjMuMzg3LjM0NC42OS43MzIuNjloNi42MjhjLjM4OCAwIC43MS0uMzAyLjczMi0uNjlsLjc0NS0xMi41MUg3LjIxek05LjcwOCA5LjI1aDEuMWwuNDU5IDguMjVoLTEuMWwtLjQ1OS04LjI1em00LjU4NCAwaC0xLjFsLS40NTkgOC4yNWgxLjFsLjQ1OS04LjI1eiIgZmlsbD0iIzRBNEE1MSIvPjxwYXRoIGQ9Ik05LjcwOCA0Ljg1aC0uMXYuMWguMXYtLjF6bTQuNTg0IDB2LjFoLjF2LS4xaC0uMXptMi44NTIgMTQuNTkybC0uMS0uMDA2LjEuMDA2em0uNzQ5LTEyLjU3NXYtLjFoLS4wOTRsLS4wMDYuMDk0LjEuMDA2em0tMi41NzkgMTQuM3YtLjEuMXptLTguNDU4LTEuNzI1bC4xLS4wMDYtLjEuMDA2ek02LjEwNyA2Ljg2N2wuMS0uMDA2LS4wMDYtLjA5NGgtLjA5NHYuMXptLTEuODk5IDBoLS4xdi4xaC4xdi0uMXptMTUuNTg0IDB2LjFoLjF2LS4xaC0uMXpNNy45NTQgMTkuMzc3bC4xLS4wMDYtLjEuMDA2ek03LjIxIDYuODY3di0uMWgtLjEwNmwuMDA2LjEwNi4xLS4wMDZ6bTguMTA0IDEzLjJ2LS4xLjF6bS43MzItLjY5bC0uMS0uMDA2LjEuMDA2em0uNzQ1LTEyLjUxbC4xLjAwNi4wMDYtLjEwNmgtLjEwNnYuMXpNMTAuODA4IDkuMjVsLjEtLjAwNi0uMDA1LS4wOTRoLS4wOTV2LjF6bS0xLjEgMHYtLjFoLS4xMDZsLjAwNi4xMDYuMS0uMDA2em0xLjU1OSA4LjI1di4xaC4xMDZsLS4wMDYtLjEwNi0uMS4wMDZ6bS0xLjEgMGwtLjEuMDA2LjAwNS4wOTRoLjA5NXYtLjF6bTMuMDI1LTguMjV2LS4xaC0uMDk1bC0uMDA1LjA5NC4xLjAwNnptMS4xIDBsLjEuMDA2LjAwNi0uMTA2aC0uMTA2di4xem0tMS41NTkgOC4yNWwtLjEtLjAwNi0uMDA2LjEwNmguMTA2di0uMXptMS4xIDB2LjFoLjA5NWwuMDA1LS4wOTQtLjEtLjAwNnptMC0xMy42NWMuMTk4IDAgLjM1OS4xNi4zNTkuMzU4aC4yYS41NTkuNTU5IDAgMDAtLjU1OS0uNTU4di4yem0tMy42NjYgMGgzLjY2NnYtLjJoLTMuNjY2di4yem0tLjM1OS4zNThjMC0uMTk3LjE2LS4zNTguMzU5LS4zNTh2LS4yYS41NTkuNTU5IDAgMDAtLjU1OS41NThoLjJ6bTAgLjY0MnYtLjY0MmgtLjJ2LjY0MmguMnptNC40ODQtLjFIOS43MDh2LjJoNC41ODR2LS4yem0tLjEtLjU0MnYuNjQyaC4ydi0uNjQyaC0uMnptMy4wNTIgMTUuMjRsLjc0OS0xMi41NzUtLjItLjAxMi0uNzQ5IDEyLjU3NS4yLjAxMnptLTEuOTMgMS44MTljMS4wMjMgMCAxLjg3LS43OTcgMS45My0xLjgybC0uMi0uMDFhMS43MzMgMS43MzMgMCAwMS0xLjczIDEuNjN2LjJ6bS02LjYyOCAwaDYuNjI4di0uMkg4LjY4NnYuMnptLTEuOTMtMS44MmExLjkzMyAxLjkzMyAwIDAwMS45MyAxLjgydi0uMmExLjczMyAxLjczMyAwIDAxLTEuNzMtMS42M2wtLjIuMDF6TTYuMDA3IDYuODc0bC43NSAxMi41NzUuMTk5LS4wMTItLjc1LTEyLjU3NS0uMTk5LjAxMnptLTEuNzk5LjA5NGgxLjg5OXYtLjJINC4yMDh2LjJ6bS0uMS0uNzQydi42NDJoLjJ2LS42NDJoLS4yem0uNTU5LS41NThhLjU1OS41NTkgMCAwMC0uNTU5LjU1OGguMmMwLS4xOTcuMTYxLS4zNTguMzU5LS4zNTh2LS4yem0xNC42NjYgMEg0LjY2N3YuMmgxNC42NjZ2LS4yem0uNTU5LjU1OGEuNTU5LjU1OSAwIDAwLS41NTktLjU1OHYuMmMuMTk4IDAgLjM1OS4xNi4zNTkuMzU4aC4yem0wIC42NDJ2LS42NDJoLS4ydi42NDJoLjJ6bS0xLjk5OS4xaDEuODk5di0uMmgtMS44OTl2LjJ6bS05Ljg0IDEyLjQwNEw3LjMxIDYuODYxbC0uMi4wMTIuNzQ0IDEyLjUxLjItLjAxMnptLjYzMy41OTZhLjYzNC42MzQgMCAwMS0uNjMyLS41OTZsLS4yLjAxMmMuMDI2LjQ0LjM5MS43ODQuODMyLjc4NHYtLjJ6bTYuNjI4IDBIOC42ODZ2LjJoNi42Mjh2LS4yem0uNjMyLS41OTZhLjYzMy42MzMgMCAwMS0uNjMyLjU5NnYuMmEuODMzLjgzMyAwIDAwLjgzMi0uNzg0bC0uMi0uMDEyem0uNzQ1LTEyLjUxbC0uNzQ1IDEyLjUxLjIuMDEyLjc0NS0xMi41MS0uMi0uMDEyem0tOS40ODEuMTA2aDkuNTgxdi0uMkg3LjIxdi4yem0zLjU5OCAyLjE4M2gtMS4xdi4yaDEuMXYtLjJ6bS41NTkgOC4zNDRsLS40Ni04LjI1LS4xOTkuMDEyLjQ2IDguMjUuMTk5LS4wMTJ6bS0xLjIuMTA2aDEuMXYtLjJoLTEuMXYuMnptLS41NTktOC4zNDRsLjQ2IDguMjUuMTk5LS4wMTItLjQ2LTguMjUtLjE5OS4wMTJ6bTMuNTg0LjA5NGgxLjF2LS4yaC0xLjF2LjJ6bS0uMzYgOC4xNTZsLjQ2LTguMjUtLjItLjAxMi0uNDU5IDguMjUuMi4wMTJ6bTEuMDAxLS4xMDZoLTEuMXYuMmgxLjF2LS4yem0uMzYtOC4xNTZsLS40NiA4LjI1LjIuMDEyLjQ1OS04LjI1LS4yLS4wMTJ6IiBmaWxsPSIjNEE0QTUxIi8+PC9zdmc+)
                    50% no-repeat;
                }

                .right-bar {
                  width: 38px;
                  height: 38px;
                  transform: rotate(90deg);
                  background: url(data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjQiIGhlaWdodD0iMjQiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PHBhdGggZD0iTTEwLjg4NiA0Ljc2NHYxNS43OGMwIC4xMjIuMDk4LjIyLjIyLjIyaDEuMDlhLjIyLjIyIDAgMDAuMjE4LS4yMlY0Ljc2NWgtMS41Mjh6IiBmaWxsPSIjNEE0QTUxIi8+PHBhdGggZD0iTTExLjk0NCAzLjE2M2w1Ljg5IDYuNDVhLjUuNSAwIDAxLS42NS43NUwxMS44NTcgNi43NGEuNS41IDAgMDAtLjU2MiAwbC01LjMzIDMuNjIzYS41LjUgMCAwMS0uNjUtLjc1bDUuODktNi40NWEuNS41IDAgMDEuNzQgMHoiIGZpbGw9IiM0QTRBNTEiLz48L3N2Zz4=)
                    50% no-repeat;
                }

                .bar-hover:hover {
                  background-color: #e3e3e4;
                  border-radius: 24px;
                  cursor: pointer;
                }
              }
            }

            .bar-padding {
              //margin-top: 10%;
              padding: 5px 8px;
            }
          }
        }
      }
    }

    .right-editor-box {
      display: block;
      margin-left: 2%;
      flex: 1;
      height: 100%;
      position: relative;

      .attribute-box {
        padding: 40px;
        background: #ffffff;
        border-radius: 8px 8px 0 0;
        width: auto;

        .custom-have-menu-box {
          display: inline-block;

          .custom-menu-name {
            display: flex;

            .input-title {
              display: inline-flex;
              float: left;
              //align-items: center;
              margin: 10px 30px 0 0;
              width: 5em;
              font-size: 14px;
            }

            .input-box {
              //display: block;

              .input-tip {
                //margin-top: 1%;
                color: #7e8081;
                font-size: 12px;
              }
            }
          }
        }

        .custom-no-have-menu-box {
          width: auto;

          .custom-menu-name {
            display: flex;
            width: auto;

            .input-title {
              display: inline-flex;
              float: left;
              //align-items: center;
              margin: 10px 30px 0 0;
              width: 5em;
              font-size: 14px;
            }
          }

          .custom-menu-type {
            display: flex;
            width: 100%;
            margin-bottom: 2%;

            .input-title {
              display: inline-flex;
              float: left;
              //align-items: center;
              margin: 10px 30px 0 0;
              width: 5em;
              font-size: 14px;
            }
          }

          .custom-menu-content {
            display: block;

            .content-select {
              display: flex;
              flex: 1;

              .send-select-box {
                background-color: #f6f7f8;
                display: inline-flex;
                padding: 6px 10px;
                white-space: nowrap;
              }

              .send-select-box::before {
                margin-right: 2px;
                display: inline-block;
                vertical-align: middle;
                float: left;
                content: ' ';
                width: 15px;
                height: 15px;
                background: transparent url('@/assets/images/sender.svg')
                  no-repeat 0 no-repeat 0 0;
                background-size: cover;
              }

              .txt-select-box {
                background-color: #f6f7f8;
                display: inline-flex;
                padding: 6px 10px;
                white-space: nowrap;
              }

              .txt-select-box::before {
                margin-right: 2px;
                display: inline-block;
                vertical-align: middle;
                float: left;
                content: ' ';
                width: 15px;
                height: 15px;
                background: transparent url('@/assets/images/sender_text.svg')
                  no-repeat 0 0;
                background-size: cover;
              }

              .img-select-box {
                background-color: #f6f7f8;
                display: inline-flex;
                padding: 6px 10px;
                white-space: nowrap;
              }

              .img-select-box::before {
                margin-right: 2px;
                display: inline-block;
                vertical-align: middle;
                float: left;
                content: ' ';
                width: 15px;
                height: 15px;
                background: transparent url('@/assets/images/sender_img.svg')
                  no-repeat 0 0;
                background-size: cover;
              }

              .radio-select-box {
                background-color: #f6f7f8;
                display: inline-flex;
                padding: 6px 10px;
                white-space: nowrap;
              }

              .radio-select-box::before {
                margin-right: 2px;
                display: inline-block;
                vertical-align: middle;
                float: left;
                content: ' ';
                width: 15px;
                height: 15px;
                background: transparent url('@/assets/images/sender_audio.svg')
                  no-repeat 0 0;
                background-size: cover;
              }

              .video-snap-select-box {
                background-color: #f6f7f8;
                display: inline-flex;
                padding: 6px 10px;
                white-space: nowrap;
              }

              .video-snap-select-box::before {
                margin-right: 2px;
                display: inline-block;
                vertical-align: middle;
                float: left;
                content: ' ';
                width: 15px;
                height: 15px;
                background: transparent url('@/assets/images/video_snap.svg')
                  no-repeat 0 0;
                background-size: cover;
              }

              .video-select-box {
                background-color: #f6f7f8;
                display: inline-flex;
                padding: 6px 10px;
                white-space: nowrap;
              }

              .video-select-box::before {
                margin-right: 2px;
                display: inline-block;
                vertical-align: middle;
                float: left;
                content: ' ';
                width: 15px;
                height: 15px;
                background: transparent url('@/assets/images/sender_video.svg')
                  no-repeat 0 0;
                background-size: cover;
              }

              .select-hover:hover {
                color: #58be6b;
                cursor: pointer;
              }
            }

            .input-box-url {
              //display: block;
            }
          }

          .custom-menu-delete {
            display: flex;
            margin-top: 30px;

            .input-title {
              flex: 1;
              display: inline-flex;
              float: left;
              //align-items: center;
              margin: 10px 30px 0 0;
              font-size: 14px;

              .delete-btn {
                display: inline-flex;
                color: #f1164e;
                cursor: pointer;
                // 禁止选中
                -webkit-user-select: none;
              }
            }
          }

          .input-box {
            //display: block;
            text-align: center;

            .input-tip {
              //margin-top: 1%;
              color: #7e8081;
              font-size: 12px;
            }
          }

          // 文字推送
          .txt-content {
            display: flex;

            .msg-container {
              display: block;
              margin: 0;
              padding: 0;
            }
          }

          // 跳转链接
          .url-content {
            display: flex;
            float: left;
            width: 100%;
            align-items: center;

            .input-box {
              //display: block;
              //text-align: center;
              .input-tip {
                //margin-top: 1%;
                color: #7e8081;
                font-size: 12px;
              }
            }
          }
        }
      }

      .menu__edit-action-bar {
        z-index: 100;
        border-top: 1px solid rgba(0, 0, 0, 0.02);
        position: sticky;
        bottom: 0;

        .menu__edit-action-inner {
          background-color: #ffffff;
          height: 80px;
          margin-left: auto;
          margin-right: auto;
          border-radius: 0 0 6px 6px;
          -webkit-box-shadow: 0 0 8px rgba(0, 0, 0, 0.02);
          box-shadow: 0 0 8px rgba(0, 0, 0, 0.02);

          .menu__edit-button {
            height: 100%;
            display: flex;
            -webkit-box-align: center;
            align-items: center;
            -webkit-box-pack: end;
            justify-content: flex-end;
            padding-right: 40px;

            .arco-btn {
              margin-right: 5px;
            }
          }
        }
      }
    }

    .attribute-box {
      height: 100%;
      -webkit-box-flex: 1;
      flex: 1;
    }
  }
</style>
