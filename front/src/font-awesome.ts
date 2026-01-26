import { library } from '@fortawesome/fontawesome-svg-core';
import {
  faArrowLeft,
  faArrowRightToBracket,
  faBookmark,
  faCalendarPlus,
  faChevronDown,
  faEdit,
  faGear,
  faMagnifyingGlass,
  faPersonCirclePlus,
  faPlus,
  faRightFromBracket,
  faStar,
  faTrash,
  faUser,
  faUserFriends,
  faWineBottle,
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

// アイコン登録
library.add(
  faArrowLeft,
  faArrowRightToBracket,
  faBookmark,
  faCalendarPlus,
  faChevronDown,
  faEdit,
  faGear,
  faMagnifyingGlass,
  faPersonCirclePlus,
  faPlus,
  faRightFromBracket,
  faStar,
  faTrash,
  faUser,
  faUserFriends,
  faWineBottle,
);

export function registerFontAwesome(
  app: ReturnType<typeof import('vue').createApp>,
) {
  app.component('font-awesome-icon', FontAwesomeIcon);
}
