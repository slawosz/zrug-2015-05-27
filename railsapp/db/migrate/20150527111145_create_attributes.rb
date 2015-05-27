class CreateAttributes < ActiveRecord::Migration
  def change
    create_table :attributes do |t|
      t.integer :product_id
      t.string :key
      t.text :value

      t.timestamps null: false
    end
  end
end
