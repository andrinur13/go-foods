-- phpMyAdmin SQL Dump
-- version 4.9.5deb2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Apr 16, 2021 at 11:33 PM
-- Server version: 8.0.23-0ubuntu0.20.04.1
-- PHP Version: 7.4.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_foods`
--

-- --------------------------------------------------------

--
-- Table structure for table `foods`
--

CREATE TABLE `foods` (
  `id` int NOT NULL,
  `name` varchar(256) NOT NULL,
  `description` text NOT NULL,
  `ingredients` text NOT NULL,
  `slug` varchar(256) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `foods`
--

INSERT INTO `foods` (`id`, `name`, `description`, `ingredients`, `slug`, `created_at`, `updated_at`) VALUES
(1, 'Kentang Pedas', 'Kentang pedas adalah salah satu khas masakan Indonesia, kentang yang dipadukan oleh sebuah sambal yang menyelimuti kentang sehingga rasa pedasnya kian terasa', 'kentang, cabe merah, cabe rawit, bawang merah, bawang putih, bawang daun, umbi kunyit, jahe, masako sapi, garam, gula merah, minyak goreng, air', 'kentang-pedas', '2021-04-16 22:59:58', '2021-04-16 22:59:58'),
(2, 'Martabak Mi Telur Ikan Kakap', 'Martabak mi adalah makanan yang sering kita jumpai di dapur kita, namun rasanya hanya seperti itu saja, kali ini yang akan membedakan martabak mie biasanya dengan martabak umumnya adalah ikan kakap yang dimasak dengan martabak mi', 'mie telur cap 3 ayam, telur, ikan kakap fillet, wortel, daun bawang, lada, garam, penyedap, minyak sayur', 'martabak-mi-telur-ikan-kakap', '2021-04-16 23:04:55', '2021-04-16 23:04:55'),
(3, 'Tempe Mendoan Krispi', 'Mendoan salah satu makanan khas Purwokerto yang memiliki bahan dasar tempe yang dilapisi oleh tepung terigu yang begitu renyah sehingga cocok untuk menemani dikala Anda sedang santai.', 'tempe, daun bawang, tepung beras, tepung terigu, baking powder, garam, air, bawang merah, bawang putih, ketumbar, garam, kunyit, kencur, kecap manis, garam, kaldu jamur', 'tempe-mendoan-krispi', '2021-04-16 23:08:04', '2021-04-16 23:08:04'),
(4, 'Butter Sugar Bread', 'Butter Sugar Bread roti yang sangat simple untuk dibuat cocok menemani di kala pandemi seperti sekarang.', 'terigu, gula pasir, garam, susu cair, telur, ragi instan, unsalted butter, salted butter, gula tabur', 'butter-sugar-bread', '2021-04-16 23:09:52', '2021-04-16 23:09:52');

-- --------------------------------------------------------

--
-- Table structure for table `food_images`
--

CREATE TABLE `food_images` (
  `id` int NOT NULL,
  `food_id` int NOT NULL,
  `is_primary` tinyint(1) NOT NULL,
  `path` varchar(512) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `food_images`
--

INSERT INTO `food_images` (`id`, `food_id`, `is_primary`, `path`, `created_at`, `updated_at`) VALUES
(1, 4, 1, 'images/20210416233123-eab14030951051e187fcc1c89a7f05b6-butter-sugar-bread.jpg', '2021-04-16 23:31:23', '2021-04-16 23:31:23'),
(2, 1, 1, 'images/20210416233153-c06e0872b9600b245eb68bf12857c386-kentang-pedas.jpg', '2021-04-16 23:31:53', '2021-04-16 23:31:53'),
(3, 2, 1, 'images/20210416233159-57ffdbae3b9fcbf0b95f83ae6b087d36-martabak-mie.jpg', '2021-04-16 23:32:00', '2021-04-16 23:32:00'),
(4, 3, 1, 'images/20210416233206-9f6f0c49d212b4b8ed48064191415670-mendoan.jpg', '2021-04-16 23:32:07', '2021-04-16 23:32:07');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `foods`
--
ALTER TABLE `foods`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `food_images`
--
ALTER TABLE `food_images`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `foods`
--
ALTER TABLE `foods`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `food_images`
--
ALTER TABLE `food_images`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
