class CreateExpenses < ActiveRecord::Migration[5.2]
  def up
    create_table :expenses do |t|
      t.datetime :paid_at
      t.integer  :amount_idr, limit: 8
      t.string   :description
      t.integer  :category_id, limit: 8
    end
  end

  def down
    drop_table :expenses
  end
end
