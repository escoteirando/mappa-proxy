<template>
  <q-page class="window-height window-width row justify-center items-center">
    <div class="column">
      <div class="row">
        <h5 class="text-h5 q-my-md">Teste Telegram</h5>
      </div>
      <div class="row">
        <q-card square bordered class="q-pa-lg shadow-1">
          <q-card-section>
            <q-form class="q-gutter-md">
              <q-input
                square
                filled
                clearable
                v-model="chatId"
                type="number"
                label="Chat Id"
              />
              <q-input
                square
                filled
                clearable
                v-model="message"
                type="text"
                label="Mensagem"
              />
            </q-form>
          </q-card-section>
          <q-card-actions class="q-px-md">
            <q-btn
              unelevated
              color="positive"
              size="lg"
              class="full-width"
              label="Enviar"
              @click="clickEnviar"
            />
          </q-card-actions>
        </q-card>
      </div>
    </div>
  </q-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { useQuasar } from 'quasar';
import { SendAuthToChat } from 'src/services/mappa';

export default defineComponent({
  name: 'TGTesting',
  components: {},
  data() {
    return {
      chatId: 0,
      message: '',
      q: useQuasar(),
    };
  },
  methods: {
    clickEnviar() {
      SendAuthToChat(this.chatId, 0, this.message)
        .then(() => {
          this.q.notify({
            caption: 'Sucesso',
            message: 'Autorização enviada para o chat no Telegram',
            icon: 'ok',
            color: 'success',
          });
        })
        .catch((error) => {
          this.q.notify({
            caption: 'Erro',
            message: `Falha no envio da autorização para o chat no Telegram (${JSON.stringify(
              error
            )})!`,
            icon: 'warning',
            color: 'negative',
          });
        });
    },
  },
});
</script>
