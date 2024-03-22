import { defineStore } from "pinia";

export interface IModalProps {
  component: null | InstanceType<typeof component>;
  props?: object;
}

export interface IModalState {
  modalState: IModalProps;
}

const basicState = { component: null, props: {} };

export default defineStore("modal-store", {
  state: (): IModalState => ({ modalState: basicState }),
  actions: {
    openModal(payload: IModalProps) {
      const { props, component } = payload;
      this.modalState = { component, props: props || {} };
    },
    closeModal() {
      this.modalState = basicState;
    },
  },
  getters: {},
});
