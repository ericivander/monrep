class CreateCategories < ActiveRecord::Migration[5.2]
  def up
    create_table :categories do |t|
      t.string  :name
      t.integer :type, limit: 2
    end
  end

  def down
    drop_table :categories
  end
end
