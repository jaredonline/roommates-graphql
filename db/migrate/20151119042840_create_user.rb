class CreateUser < ActiveRecord::Migration
  def change
    create_table "debts", :force => true do |t|
      t.integer  "amount_in_cents"
      t.integer  "person_id"
      t.integer  "loaner_id"
      t.integer  "expense_id"
      t.boolean  "paid"
      t.datetime "created_at"
      t.datetime "updated_at"
    end

    add_index "debts", ["expense_id"], :name => "index_debts_on_expense_id"
    add_index "debts", ["loaner_id"], :name => "index_debts_on_loaner_id"
    add_index "debts", ["person_id"], :name => "index_debts_on_person_id"

    create_table "delayed_jobs", :force => true do |t|
      t.integer  "priority",   :default => 0
      t.integer  "attempts",   :default => 0
      t.text     "handler"
      t.text     "last_error"
      t.datetime "run_at"
      t.datetime "locked_at"
      t.datetime "failed_at"
      t.text     "locked_by"
      t.datetime "created_at"
      t.datetime "updated_at"
    end

    create_table "expenses", :force => true do |t|
      t.integer  "amount_in_cents"
      t.integer  "house_id"
      t.string   "name"
      t.text     "description"
      t.datetime "created_at"
      t.datetime "updated_at"
      t.integer  "creator_id"
      t.text     "notes"
      t.integer  "loaner_id"
    end

    add_index "expenses", ["creator_id"], :name => "index_expenses_on_creator_id"
    add_index "expenses", ["house_id"], :name => "index_expenses_on_house_id"
    add_index "expenses", ["loaner_id"], :name => "index_expenses_on_loaner_id"

    create_table "expenses_people", :id => false, :force => true do |t|
      t.integer "expense_id"
      t.integer "person_id"
    end

    add_index "expenses_people", ["expense_id"], :name => "index_expenses_people_on_expense_id"
    add_index "expenses_people", ["person_id"], :name => "index_expenses_people_on_person_id"

    create_table "houses", :force => true do |t|
      t.string   "name"
      t.datetime "created_at"
      t.datetime "updated_at"
    end

    create_table "payments", :force => true do |t|
      t.integer  "person_paid_id"
      t.integer  "person_paying_id"
      t.integer  "amount_in_cents"
      t.datetime "created_at"
      t.datetime "updated_at"
      t.text     "notes"
      t.integer  "house_id"
    end

    add_index "payments", ["house_id"], :name => "index_payments_on_house_id"
    add_index "payments", ["person_paid_id"], :name => "index_payments_on_person_paid_id"
    add_index "payments", ["person_paying_id"], :name => "index_payments_on_person_paying_id"

    create_table "people", :force => true do |t|
      t.string   "name"
      t.string   "email"
      t.datetime "created_at"
      t.datetime "updated_at"
      t.integer  "house_id"
      t.string   "crypted_password",          :limit => 128, :default => "", :null => false
      t.string   "salt",                      :limit => 128, :default => "", :null => false
      t.string   "remember_token",            :limit => 40
      t.datetime "remember_token_expires_at"
      t.string   "persistence_token",                        :default => "", :null => false
      t.string   "perishable_token",                         :default => "", :null => false
      t.integer  "login_count",                              :default => 0,  :null => false
      t.integer  "failed_login_count",                       :default => 0,  :null => false
      t.datetime "last_request_at"
      t.datetime "current_login_at"
      t.datetime "last_login_at"
      t.string   "current_login_ip"
      t.string   "last_login_ip"
    end

    add_index "people", ["house_id"], :name => "index_people_on_house_id"

    create_table "signups", :force => true do |t|
      t.string   "email"
      t.string   "name"
      t.datetime "created_at"
      t.datetime "updated_at"
    end
  end
end
