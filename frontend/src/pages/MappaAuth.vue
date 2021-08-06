<template>
  <q-page class="window-height window-width row justify-center items-center">
    <div class="column">
      <div class="row">
        <h5 class="text-h5 text-white q-my-md">Autenticação mAPPa</h5>
      </div>
      <div class="row">
        <q-card square bordered class="q-pa-lg shadow-1" v-if="!authKey">
          <q-card-section>
            <q-form class="q-gutter-md">
              <q-input
                square
                filled
                clearable
                v-model="username"
                type="text"
                label="Usuário mAPPa"
              />
              <q-input
                square
                filled
                clearable
                v-model="password"
                type="password"
                label="Senha"
              />
            </q-form>
          </q-card-section>
          <q-card-actions class="q-px-md">
            <q-btn
              unelevated
              color="positive"
              size="lg"
              class="full-width"
              label="Login"
              @click="clickLogin"
            />
          </q-card-actions>
          <q-card-section class="text-center q-pa-none">
            <p class="text-grey-16">
              Sua senha não será armazenada.<br />Esta autenticação é usada para
              se obter acesso aos dados do mAPPa.
            </p>
          </q-card-section>
        </q-card>
        <q-card square bordered class="q-pa-lg shadow-1" v-else>
          <q-card-section>
            <div class="text-h6">Chave de autorização</div>
          </q-card-section>
          <q-separator inset />
          <q-card-section>
            <q-input
              square
              filled
              clearable
              v-model="authKey"
              type="text"
              label="Chave"
              readonly
            />
          </q-card-section>
          <q-card-actions class="q-px-md">
            <q-btn
              unelevated
              color="positive"
              size="lg"
              class="full-width"
              label="Copiar"
              @click="clickCopy"
              icon="content_copy"
            />
          </q-card-actions>
          <q-card-section class="text-center q-pa-none">
            <p class="text-grey-16">
              Esta chave será utilizada para autorização no grupo da sua seção
              no Telegram.<br />
              Copie e cole no grupo para que o nosso robô faça a conexão
              corretamente.
            </p>
          </q-card-section>
        </q-card>
      </div>
    </div>
  </q-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { useQuasar } from 'quasar';
import { ParseContext, EmptyLoginContext, DoLogin } from 'src/services/mappa';

export default defineComponent({
  name: 'MappaAuth',
  components: {},
  data() {
    return {
      username: '',
      password: '',
      authKey: '',
      context: EmptyLoginContext,
      clipboardEnabled: false,
      q: useQuasar(),
    };
  },
  mounted() {
    this.context = ParseContext(this.$route.params.context as string);
  },
  methods: {
    clickCopy() {
      navigator.clipboard.writeText(this.authKey).then(
        () => {
          this.q.notify({
            caption: 'Área de transferência',
            message:
              'Chave foi copiada com sucesso. Você pode colá-la no grupo do telegram para continuar o setup.',
            icon: 'check_circle',
            color: 'success',
          });
        },
        () => {
          this.q.notify({
            caption: 'Área de transferência',
            message:
              'Erro na cópia para área de transferência. Faça o processo manualmente.',
            icon: 'warning',
            color: 'negative',
          });
        }
      );
    },
    async clickLogin() {
      if (!this.username || !this.password) {
        this.q.notify({
          caption: 'Erro',
          message: 'Informe usuário e senha para continuar!',
          icon: 'warning',
          color: 'negative',
        });
        return;
      }

      try {
        this.q.notify({
          caption: 'Conectando',
          message: 'Login no proxy mAPPa',
          icon: 'link',
          group: false,
          color: 'primary',
        });

        let loginResponse = await DoLogin(this.username, this.password);
        if (!loginResponse) {
          this.q.notify({
            caption: 'Erro',
            message: 'Falha na conexão com o proxy mAPPa!',
            icon: 'warning',
            color: 'negative',
          });
          return;
        }
        let responseKey = JSON.stringify(loginResponse);
        this.authKey = `/auth ${btoa(responseKey)}`;
        // this.q.notify({
        //   caption: 'Resposta',
        //   message: responseKey,
        //   icon: 'check_circle',
        //   color: 'success',
        //   group: false,
        // });
      } catch (err) {
        console.error('Login error:', err);
        this.q.notify({
          caption: 'Erro',
          message: (err as Error).message,
          icon: 'danger',
          group: false,
        });
      }
    },
  },
});
</script>
