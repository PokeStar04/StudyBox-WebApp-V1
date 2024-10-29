import { useNavigate } from 'react-router-dom';
import { useSelector } from 'react-redux';
import Breadcrumb from '../Breadcrumb/Breadcrumb';
import { RootState } from '../../../../../store';

interface TopBarProps {
  isCollapsed: boolean;
}

const TopBar = ({ isCollapsed }: TopBarProps) => {
  const navigate = useNavigate();

  // Récupérer l'image de profil et les autres infos depuis le store Redux
  const profileImage = useSelector(
    (state: RootState) => state.profileImage.profileImage,
  );

  const defaultImageUrl = ''; // Image par défaut

  const handleProfileClick = () => {
    navigate('/profil');
  };

  // Fonction de gestion d'erreur de chargement de l'image
  const handleImageError = (
    event: React.SyntheticEvent<HTMLImageElement, Event>,
  ) => {
    event.currentTarget.src = defaultImageUrl; // Remplace l'image par défaut en cas d'erreur
  };

  return (
    <div
      className={`fixed top-0 left-0 right-0 z-10 bg-white px-8 py-4 border-b border-lightGray flex justify-between items-center z-50 ${
        isCollapsed ? 'ml-32' : 'ml-72'
      }`}
    >
      {/* Breadcrumb à gauche */}
      <div className="flex-1">
        <Breadcrumb />
      </div>

      {/* Image de profil à droite */}
      <div className="flex items-center">
        <img
          src={profileImage || defaultImageUrl}
          alt="Profile"
          className="w-10 h-10 rounded-full border border-gray-300 cursor-pointer"
          onClick={handleProfileClick}
          onError={handleImageError}
        />
      </div>
    </div>
  );
};

export default TopBar;
