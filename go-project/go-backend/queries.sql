CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    role ENUM('admin', 'user') DEFAULT 'user', -- Role for authorization (e.g., admin or regular user)
    is_active BOOLEAN DEFAULT TRUE,           -- Whether the user account is active
    reset_token VARCHAR(255),                 -- Token for password reset
    api_key VARCHAR(64) UNIQUE,                -- API key for accessing the API
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE users_test (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    role ENUM('admin', 'user') DEFAULT 'user', -- Role for authorization (e.g., admin or regular user)
    is_active BOOLEAN DEFAULT TRUE,           -- Whether the user account is active
    reset_token VARCHAR(255),                 -- Token for password reset
    api_key VARCHAR(64) UNIQUE,                -- API key for accessing the API
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);