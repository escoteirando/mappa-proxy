<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn flat dense round icon="menu" aria-label="Menu" />

        <q-toolbar-title :title="lastUpdate">
          Mappa Frontend {{ version }}</q-toolbar-title
        >

        <div>
          <q-btn
            :title="proxyHCTitle"
            flat
            dense
            :icon="proxyIcon"
            :label="proxyHC"
          />
        </div>
      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { DoHC } from 'src/services/mappa';
import { version, lastUpdate } from 'src/assets/version';

export default defineComponent({
  name: 'SingleLayout',

  components: {},
  data() {
    return {
      proxyHC: '...',
      proxyHCTitle: '',
      proxyIcon: 'link',
      version: `v${version.major}.${version.minor}.${version.feature}`,
      lastUpdate: `Último build: ${lastUpdate.toLocaleString()}`,
    };
  },

  async mounted() {
    try {
      let hc = await DoHC();
      console.log('HC', hc);
      this.proxyHCTitle = `Server: ${hc.mappa_server.status}`;
      this.proxyHC = hc.status == 'HEALTHY' ? 'OK' : 'Indisponível';
      this.proxyIcon = this.proxyHC == 'OK' ? 'link' : 'link_off';
    } catch (err) {
      console.error('HC', err);
      this.proxyHCTitle = (err as Error).message;
      this.proxyHC = 'Em erro';
      this.proxyIcon = 'link_off';
    }
    this.proxyHC = `Proxy: ${this.proxyHC}`;
  },
});
</script>
