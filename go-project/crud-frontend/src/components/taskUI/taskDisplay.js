import React, { useState } from "react";

export default function TaskCard({id, title, content, onEdit, onDelete }) {
  // State to toggle edit mode
  const [isEditing, setIsEditing] = useState(false);
  const [editedTitle, setEditedTitle] = useState(title);
  const [editedContent, setEditedContent] = useState(content);

  const handleSave = () => {
    onEdit({id, title: editedTitle, content: editedContent });
    setIsEditing(false);
  };

  return (
    <div className="bg-white shadow-md rounded-lg p-4 w-full mx-auto border border-gray-200">
      {isEditing ? (
        <>
          {/* Editable Title */}
          <input
            type="text"
            value={editedTitle}
            onChange={(e) => setEditedTitle(e.target.value)}
            className="w-full p-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          />

          {/* Editable Content */}
          <textarea
            value={editedContent}
            onChange={(e) => setEditedContent(e.target.value)}
            className="w-full mt-2 p-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
            rows="3"
          />

          {/* Save and Cancel Buttons */}
          <div className="flex justify-end mt-4 space-x-3">
            <button
              onClick={handleSave}
              className="px-4 py-2 text-sm font-medium text-white bg-green-500 rounded-lg hover:bg-green-600"
            >
              Save
            </button>
            <button
              onClick={() => {
                if (window.confirm("Are you sure you want to cancel?")) {
                  setIsEditing(false); //Exit edit mode
                  setEditedTitle(title); //Revert title to its original
                  setEditedContent(content); //Revert content to its original
                }
              }}
              className="px-4 py-2 text-sm font-medium text-white bg-gray-500 rounded-lg hover:bg-gray-600"
            >
              Cancel
            </button>
          </div>
        </>
      ) : (
        <>
          {/* Display Title */}
          <h3 className="text-lg font-semibold text-gray-800">{title}</h3>

          {/* Display Content */}
          <p className="text-gray-600 mt-2">{content}</p>

          {/* Action Buttons */}
          <div className="flex justify-end mt-4 space-x-3">
            <button
              onClick={() => setIsEditing(true)}
              className="px-4 py-2 text-sm font-medium text-white bg-blue-500 rounded-lg hover:bg-blue-600"
            >
              Edit
            </button>
            <button
              onClick= {onDelete}
              className="px-4 py-2 text-sm font-medium text-white bg-red-500 rounded-lg hover:bg-red-600"
            >
              Delete
            </button>
          </div>
        </>
      )}
    </div>
  );
}
