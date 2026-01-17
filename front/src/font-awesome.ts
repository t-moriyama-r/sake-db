import { library } from '@fortawesome/fontawesome-svg-core';
import {
  faArrowLeft,
  faArrowRightToBracket,
  faMagnifyingGlass,
  faPersonCirclePlus,
  faPlus,
  faRightFromBracket,
  faWineBottle,
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

// アイコン登録
library.add(
  faArrowLeft,
  faArrowRightToBracket,
  faMagnifyingGlass,
  faPersonCirclePlus,
  faPlus,
  faRightFromBracket,
  faWineBottle,
);

export function registerFontAwesome(
  app: ReturnType<typeof import('vue').createApp>,
) {
  app.component('font-awesome-icon', FontAwesomeIcon);
}
