CREATE TABLE `roles` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NOT NULL,
  `created_at` bigint(20) NOT NULL,
  `createdBy` bigint(20) NOT NULL
);

CREATE TABLE `users` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(255),
  `password` VARCHAR(50) NOT NULL,
  `role_id` bigint(20),
  `updated_at` bigint(20) NOT NULL,
  `created_at` bigint(20) NOT NULL
);

CREATE TABLE `user_details` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20),
  `email` VARCHAR(255),
  `name` VARCHAR(255),
  `date_of_birth` VARCHAR(255),
  `updated_at` bigint(20) NOT NULL,
  `created_at` bigint(20) NOT NULL
);

CREATE TABLE `user_education_details` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20),
  `certificate_degree_name` VARCHAR(255),
  `major` VARCHAR(255),
  `institute_university_name` VARCHAR(255),
  `starting_date` VARCHAR(255),
  `end_date` VARCHAR(255),
  `percentage` bigint(20),
  `cgpa` bigint(20),
  `updated_at` bigint(20) NOT NULL,
  `created_at` bigint(20) NOT NULL
);

CREATE TABLE `user_experience_details` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20),
  `is_current_job` VARCHAR(255),
  `start_date` VARCHAR(255),
  `end_date` VARCHAR(255),
  `company_name` VARCHAR(255),
  `job_location_city` VARCHAR(255),
  `job_location_state` VARCHAR(255),
  `job_location_country` VARCHAR(255),
  `updated_at` bigint(20) NOT NULL,
  `created_at` bigint(20) NOT NULL
);

CREATE TABLE `attachments` (
  `id` bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `title` varchar(256),
  `description` longtext,
  `name` varchar(256),
  `size` int,
  `type` varchar(60),
  `status` int(1),
  `updated_at` bigint(20) NOT NULL,
  `created_at` bigint(20) NOT NULL
);

INSERT INTO `users`(`id`, `email`,  `password`, `role_id`, `updated_at`, `created_at`) VALUES (1, 'line.manager@gmail.com', '8a55660ad87534091fce106237172b4c21aebc62', 1, 123213, 12312312);
INSERT INTO `user_details`(`id`, `user_id`, `email`, `name`, `date_of_birth`, `updated_at`, `created_at`) VALUES (1, 1, 'line.manager@gmail.com', 'Mr Awsome Manager', "1989-01-01", 123213, 12312312);