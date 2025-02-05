import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

type ModalProps = {
  isOpen: boolean;
  onClose: () => void;
  status: "loading" | "success" | "error";
  isUpdate: boolean; // ✅ Ajout pour différencier création et mise à jour
};

export default function EventModal({ isOpen, onClose, status, isUpdate }: ModalProps) {
  const navigate = useNavigate();

  useEffect(() => {
    if (!isOpen) return;
    const handleEsc = (event: KeyboardEvent) => {
      if (event.key === "Escape") handleClose();
    };
    window.addEventListener("keydown", handleEsc);
    return () => window.removeEventListener("keydown", handleEsc);
  }, [isOpen]);

  if (!isOpen) return null;

  const handleClose = () => {
    onClose();
    navigate("/"); // ✅ Redirige vers /profil après fermeture
  };

  // ✅ Détermine le texte affiché en fonction de isUpdate
  const title =
    status === "loading"
      ? isUpdate
        ? "Événement en cours de mise à jour..."
        : "Événement en cours de création..."
      : status === "success"
      ? isUpdate
        ? "Événement mis à jour avec succès"
        : "Événement créé avec succès"
      : "Échec de l'opération";

  const description =
    status === "success"
      ? isUpdate
        ? "L'événement a été mis à jour avec succès !"
        : "L'événement a été ajouté avec succès !"
      : "Une erreur s'est produite, veuillez réessayer.";

  return (
    <div
      className="fixed inset-0 flex items-center justify-center bg-gray-800 bg-opacity-50 transition-opacity duration-300 animate-fadeIn"
      onClick={handleClose}
    >
      <div
        className="bg-white rounded-xl shadow-lg p-6 max-w-sm text-center animate-scaleIn"
        onClick={(e) => e.stopPropagation()}
        role="dialog"
        aria-labelledby="modal-title"
      >
        {status === "loading" ? (
          <>
            <div className="flex justify-center items-center mb-4">
              <div className="w-12 h-12 border-4 border-blue-500 border-t-transparent rounded-full animate-spin"></div>
            </div>
          </>
        ) : (
          <>
            <div className="flex justify-center items-center mb-4">
              <div
                className={`w-12 h-12 flex items-center justify-center rounded-full ${
                  status === "success" ? "bg-green-100 animate-bounceIn" : "bg-red-100 animate-shake"
                }`}
              >
                {status === "success" ? (
                  <span className="text-green-500 text-3xl">✓</span>
                ) : (
                  <span className="text-red-500 text-3xl">✗</span>
                )}
              </div>
            </div>
          </>
        )}

        <h2 id="modal-title" className="text-lg font-semibold">{title}</h2>
        <p className="text-gray-500 text-sm mt-2">{description}</p>

        <button
          onClick={handleClose}
          className="mt-4 bg-blue-600 text-white px-4 py-2 rounded-lg w-full hover:bg-blue-700 transition"
        >
          Revenir à l'accueil
        </button>
      </div>
    </div>
  );
}
