'use strict';

module.exports = {
  up: async (queryInterface, Sequelize) => {
    return queryInterface.createTable('artists', {
        id: {
          type: Sequelize.BIGINT,
          allowNull: false,
          primaryKey: true,
          autoIncrement: true
      },
      artist_name:{
          type: Sequelize.STRING,
          allowNull: false,
      },
      album_name:{
        type: Sequelize.STRING,
        allowNull: false,
      },
      image_url: {
        type: Sequelize.TEXT,
        allowNull: true
      },
      release_date: {
        type: Sequelize.DATE,
        allowNull: true
      },
      price: {
        type: Sequelize.DECIMAL(10,2),
        allowNull: false
      },
      sample_url: {
        type: Sequelize.TEXT,
        allowNull: true
      },
      created_at: {
        type: Sequelize.DATE,
        allowNull: false
      },
      updated_at: {
          type: Sequelize.DATE,
          allowNull: false
      },
      deleted_at: {
          type: Sequelize.DATE,
          allowNull: true
      }
    })
  },

  down: async (queryInterface, Sequelize) => {
    return queryInterface.dropTable('artists');
  }
};
