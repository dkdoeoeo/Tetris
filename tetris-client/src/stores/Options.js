import { defineStore } from "pinia";
const defaultSettings = {
    MusicVolume: 15,
    Theme: 0,
}
const getSettings = () => {
    const settings = localStorage.getItem("OptionsStore")
  
    return settings ? JSON.parse(settings) : defaultSettings
}

export const useOptionsStore = defineStore("OptionsStore", {
    state: () => ({
      settings: defaultSettings
    }),
    actions: {
      updateSettings(partialSettings) {
        this.settings = {
          ...this.settings,
          ...partialSettings,
        };
  
        localStorage.setItem("OptionsStore", JSON.stringify(this.settings));
      },
    },
});
  