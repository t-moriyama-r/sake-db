import { library } from '@fortawesome/fontawesome-svg-core';
import {
  faArrowRightToBracket,
  faPersonCirclePlus,
  faPlus,
  faWineBottle,
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

// アイコン登録
library.add(faArrowRightToBracket, faPersonCirclePlus, faPlus, faWineBottle);

export function registerFontAwesome(
  app: ReturnType<typeof import('vue').createApp>,
) {
  app.component('font-awesome-icon', FontAwesomeIcon);
}
