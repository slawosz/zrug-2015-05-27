class Product < ActiveRecord::Base
  has_many :attrs, table_name: :attributes, class: Attribute
end
